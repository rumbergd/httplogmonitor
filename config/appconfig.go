package config

// InputFile is path to the input http access log file
var InputFile string

// StatsConfig is configuration for collecting stats from http log
var StatsConfig StatsConfiguration

// AlertsConfig is configuration for creating alerts
var AlertsConfig AlertsConfiguration

// StatsConfiguration is a struct representing configuration for collecting statistics from the http access log
type StatsConfiguration struct {
	IntervalSeconds int
}

// AlertsConfiguration is a struct representing configuration for different alerts. Currently there is only VolumeAlert
type AlertsConfiguration struct {
	VolumeAlert VolumeAlertConfiguration
}

// VolumeAlertConfiguration is a struct representing configuration for volume alerts
type VolumeAlertConfiguration struct {
	StatsIntervalLookback int
	Threshold             int64
	LogFile               string
}
