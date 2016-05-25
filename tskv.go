package tskv

import (
	"fmt"
	"time"
)

// Tskv our main structure
type Tskv struct {
	duration  time.Duration
	frequency time.Duration
	elements  map[time.Time]Element
}

// Config holds our main configuration information
type Config struct {
	Duration  time.Duration
	Frequency time.Duration
}

// New Returns a brand new Tskv object
func New(config *Config) (*Tskv, error) {
	duration := config.Duration
	if duration == 0 {
		return nil, fmt.Errorf("Duration must be >0")
	}
	frequency := config.Frequency
	if frequency == 0 {
		return nil, fmt.Errorf("Duration must be >0")
	}
	if frequency > duration {
		return nil, fmt.Errorf("Duration must be larger than Frequency")
	}

	t := &Tskv{
		duration:  duration,
		frequency: frequency,
		elements:  make(map[time.Time]Element),
	}
	return t, nil
}
