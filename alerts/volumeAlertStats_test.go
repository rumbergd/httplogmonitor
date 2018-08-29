package alerts

import (
	"github.com/oleiade/lane"
	"github.com/stretchr/testify/assert"
	"httplogmonitor/config"
	"log"
	"os"
	"testing"
)

func TestVolumeAlertStats_Add(t *testing.T) {
	actual := VolumeAlertStats{
		totalHits:       int64(60),
		isVolumeAlertON: false,
		hitsQueue:       lane.NewQueue(),
	}
	actual.hitsQueue.Enqueue(int64(10))
	actual.hitsQueue.Enqueue(int64(20))
	actual.hitsQueue.Enqueue(int64(30))

	newHits := int64(40)
	expected := VolumeAlertStats{
		totalHits:       int64(100),
		isVolumeAlertON: false,
		hitsQueue:       lane.NewQueue(),
	}
	expected.hitsQueue.Enqueue(int64(10))
	expected.hitsQueue.Enqueue(int64(20))
	expected.hitsQueue.Enqueue(int64(30))
	expected.hitsQueue.Enqueue(int64(40))

	actual.Add(newHits)

	assert.Equal(t, actual.totalHits, expected.totalHits, "Add() provided unexpected result")
	assert.Equal(t, actual.isVolumeAlertON, expected.isVolumeAlertON, "Add() provided unexpected result")
	assert.Equal(t, *actual.hitsQueue, *expected.hitsQueue, "Add() provided unexpected result")
}

func TestVolumeAlertStats_AddWithQueueLimit(t *testing.T) {
	config.AlertsConfig.VolumeAlert.StatsIntervalLookback = 5

	actual := VolumeAlertStats{
		totalHits:       int64(150),
		isVolumeAlertON: false,
		hitsQueue:       lane.NewQueue(),
	}
	actual.hitsQueue.Enqueue(int64(10))
	actual.hitsQueue.Enqueue(int64(20))
	actual.hitsQueue.Enqueue(int64(30))
	actual.hitsQueue.Enqueue(int64(40))
	actual.hitsQueue.Enqueue(int64(50))

	newHits := int64(60)
	expected := VolumeAlertStats{
		totalHits:       int64(200),
		isVolumeAlertON: false,
		hitsQueue:       lane.NewQueue(),
	}
	expected.hitsQueue.Enqueue(int64(20))
	expected.hitsQueue.Enqueue(int64(30))
	expected.hitsQueue.Enqueue(int64(40))
	expected.hitsQueue.Enqueue(int64(50))
	expected.hitsQueue.Enqueue(int64(60))

	actual.Add(newHits)

	assert.Equal(t, actual.totalHits, expected.totalHits, "Add() provided unexpected result")
	assert.Equal(t, actual.isVolumeAlertON, expected.isVolumeAlertON, "Add() provided unexpected result")
	assert.Equal(t, *actual.hitsQueue, *expected.hitsQueue, "Add() provided unexpected result")
}

func TestVolumeAlertStats_EvaluateThresholdCreateAlert(t *testing.T) {
	config.AlertsConfig.VolumeAlert.Threshold = 100
	config.AlertsConfig.VolumeAlert.StatsIntervalLookback = 5
	volumeAlertLogger = log.New(os.Stdout, "", 0)

	actual := VolumeAlertStats{
		totalHits:       int64(150),
		isVolumeAlertON: false,
		hitsQueue:       lane.NewQueue(),
	}
	actual.hitsQueue.Enqueue(int64(10))
	actual.hitsQueue.Enqueue(int64(20))
	actual.hitsQueue.Enqueue(int64(30))
	actual.hitsQueue.Enqueue(int64(40))
	actual.hitsQueue.Enqueue(int64(50))

	actual.EvaluateThreshold()

	assert.Equal(t, actual.isVolumeAlertON, true, "EvaluateThreshold() provided unexpected result")
}

func TestVolumeAlertStats_EvaluateThresholdRecoverAlert(t *testing.T) {
	config.AlertsConfig.VolumeAlert.Threshold = 100
	config.AlertsConfig.VolumeAlert.StatsIntervalLookback = 5
	volumeAlertLogger = log.New(os.Stdout, "", 0)

	actual := VolumeAlertStats{
		totalHits:       int64(90),
		isVolumeAlertON: true,
		hitsQueue:       lane.NewQueue(),
	}
	actual.hitsQueue.Enqueue(int64(10))
	actual.hitsQueue.Enqueue(int64(20))
	actual.hitsQueue.Enqueue(int64(30))
	actual.hitsQueue.Enqueue(int64(40))
	actual.hitsQueue.Enqueue(int64(0))

	actual.EvaluateThreshold()

	assert.Equal(t, actual.isVolumeAlertON, false, "EvaluateThreshold() provided unexpected result")
}
