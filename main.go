package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"httplogmonitor/alerts"
	"httplogmonitor/config"
	"httplogmonitor/stats"
	"os"
	"time"
)

func init() {
	loadConfigurations("appconfig.yaml")
}

func main() {

	// Error out if no input file provided
	if config.InputFile == "" {
		fmt.Printf("INPUTFILE is not set appconfig.yaml")
		os.Exit(-1)
	}

	// Setup logger for http stats
	stats.SetupLogger()

	// Setup logger for volume alerts
	alerts.SetupLogger(config.AlertsConfig.VolumeAlert.LogFile)

	// Setup ticker to print http access log stats every IntervalSeconds (10 seconds by default)
	ticker := time.NewTicker(time.Second * time.Duration(config.StatsConfig.IntervalSeconds))
	go updateLogStats(ticker)

	// Start tailing input http access log file
	t, err := tail.TailFile(config.InputFile, tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}

	// Process lines as they're appended to the input file
	for line := range t.Lines {
		// Update http traffic statistics for every log line
		stats.LogStats.Add(stats.CreateHTTPLogEntry(line.Text))
	}
}

// updateLogStats is invoked by timer every IntervalSeconds (10 sec by default).
// It updates the log stats and checks against high volume threshold to generate an alert.
func updateLogStats(ticker *time.Ticker) {
	for range ticker.C {
		// Print the log stats for the last timer interval
		stats.LogStats.Log()

		// Add nuber of hits for the last timer interval to the alert stats collection
		alerts.AlertStats.Add(stats.LogStats.GetHits())
		// Evaluate if number of hits crossed the threshold and create the alert if it did
		alerts.AlertStats.EvaluateThreshold()

		// Reset the log stats for the next timer interval period
		stats.LogStats.Reset()
	}
}
