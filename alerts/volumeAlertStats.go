package alerts

import (
	"github.com/oleiade/lane"
	"httplogmonitor/config"
	"sync"
	"time"
)

// AlertStats represents statistics about http traffic collected for evaluating alert thresholds
var AlertStats VolumeAlertStats

// VolumeAlertStats is a type that has data for evaluating volume alert thresholds
type VolumeAlertStats struct {
	sync.RWMutex
	hitsQueue       *lane.Queue
	totalHits       int64
	isVolumeAlertON bool
}

// Add function adds new value of hits collected every 10 sec to the queue and updates total hits count
func (s *VolumeAlertStats) Add(hits int64) {
	// Ensure the access is thread safe
	s.Lock()
	defer s.Unlock()

	// Create queue if it doesn't exist
	if s.hitsQueue == nil {
		s.hitsQueue = lane.NewQueue()
	}

	// Limit size of the queue to only monitor last several stats collection intervals. Default value is 2 minutes (12*10sec = 2min)
	if s.hitsQueue.Size() == config.AlertsConfig.VolumeAlert.StatsIntervalLookback {
		// Evict oldest statistic from the queue
		removedHits := s.hitsQueue.Dequeue()
		s.totalHits -= removedHits.(int64)
	}

	// Add new statistic to the queue
	s.hitsQueue.Enqueue(hits)
	s.totalHits += hits
}

// EvaluateThreshold checks totalHits against the threshold and creates alert
func (s *VolumeAlertStats) EvaluateThreshold() {
	// Create and log alert if threshold is crossed and alert is not ON already
	// Also, hitsQueue must be filled up to ensure the alerts aren't created prematurely when the monitor starts
	if s.totalHits > config.AlertsConfig.VolumeAlert.Threshold &&
		!s.isVolumeAlertON &&
		s.hitsQueue.Size() == config.AlertsConfig.VolumeAlert.StatsIntervalLookback {

		a := CreateVolumeAlert(s.totalHits)
		s.isVolumeAlertON = true
		volumeAlertLogger.Printf("[%v] %v\n", time.Now(), a.ToString())
	}
	// Recover alert if totalHits is back below threshold and alert was ON
	if s.totalHits <= config.AlertsConfig.VolumeAlert.Threshold && s.isVolumeAlertON {
		s.isVolumeAlertON = false
		volumeAlertLogger.Printf("[%v] High traffic alert recovered\n", time.Now())
	}
}
