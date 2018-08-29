package alerts

import (
	"fmt"
	"time"
)

// VolumeAlert is a type that represent an alert generated for high traffic
type VolumeAlert struct {
	timestamp time.Time
	hits      int64
}

// ToString returns string representation of the volume alert
func (a *VolumeAlert) ToString() string {
	return fmt.Sprintf("High traffic generated an alert - hits = %v, triggered at %v",
		a.hits, a.timestamp,
	)
}

// CreateVolumeAlert creates volume alert with passed number of http hits
func CreateVolumeAlert(hits int64) VolumeAlert {
	return VolumeAlert{
		timestamp: time.Now(),
		hits:      hits,
	}
}
