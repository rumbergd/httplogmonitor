package alerts

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// volumeAlertLogger is used for recording/saving/printing alerts
var volumeAlertLogger *log.Logger

// SetupLogger initializes alert logger to print to console and save alerts to a file
func SetupLogger(filename string) {

	// Open file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// Create logger
	volumeAlertLogger = log.New(io.MultiWriter(file, os.Stdout), "", 0)

	// Close log file on program exit
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		file.Close()
		os.Exit(0)
	}()
}
