package alerts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVolumeAlert_ToString(t *testing.T) {
	ts := time.Now()
	a := VolumeAlert{
		timestamp: ts,
		hits:      1234,
	}
	expected := fmt.Sprintf("High traffic generated an alert - hits = 1234, triggered at %v", ts)
	actual := a.ToString()
	assert.Equal(t, actual, expected, "ToString() provided unexpected result")
}

func TestVolumeAlert_CreateVolumeAlert(t *testing.T) {
	hits := int64(1234)
	expected := VolumeAlert{
		hits: int64(1234),
	}

	actual := CreateVolumeAlert(hits)
	assert.Equal(t, actual.hits, expected.hits, "CreateVolumeAlert() provided unexpected result for clientIP field")
}
