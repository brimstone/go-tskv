package tskv

import "time"

// Tskv our main structure
type Tskv struct {
	Duration  time.Time
	Frequency time.Duration
	elements  map[string]interface{}
}

// AddNow Adds an element to our tskv, for right now
func (t *Tskv) Add(i interface{}) {
}
