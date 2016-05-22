package tskv

import "time"

// Tskv our main structure
type Tskv struct {
	Duration  time.Time
	Frequency time.Duration
	elements  map[time.Time]Element
}

// New Returns a brand new Tskv object
func New() (*Tskv, error) {
	t := &Tskv{
		elements: make(map[time.Time]Element),
	}
	return t, nil
}
