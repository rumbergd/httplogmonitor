package main

import (
	"github.com/spf13/viper"
	"httplogmonitor/config"
)

// loadConfigurations reads configFileNamePath and loads all the configuration parameters
func loadConfigurations(configFileNamePath string) {

	v := viper.New()
	v.SetConfigFile(configFileNamePath)
	v.ReadInConfig()

	config.InputFile = v.GetString("inputfile")

	config.StatsConfig = config.StatsConfiguration{
		IntervalSeconds: v.GetInt("stats.intervalSeconds"),
	}

	config.AlertsConfig = config.AlertsConfiguration{
		VolumeAlert: config.VolumeAlertConfiguration{
			StatsIntervalLookback: v.GetInt("alerts.volumeAlert.statsIntervalLookback"),
			Threshold:             v.GetInt64("alerts.volumeAlert.threshold"),
			LogFile:               v.GetString("alerts.volumeAlert.logFile"),
		},
	}
}
