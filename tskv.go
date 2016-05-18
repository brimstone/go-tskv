package tskv

import "time"

// Tskv our main structure
type Tskv struct {
	Duration  time.Time
	Frequency time.Duration
	elements  map[string]interface{}
}

// New Returns a brand new Tskv object
func New() (*Tskv, error) {
	t := &Tskv{
		elements: make(map[string]interface{}),
	}
	return t, nil
}
