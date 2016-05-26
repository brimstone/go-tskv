package tskv

import "time"

// Set Sets an element to our tskv
func (t *Tskv) Set(when time.Time, value interface{}) error {
	t.cleanup()
	t.elements[when.Truncate(t.frequency)] = Element{value: value}
	return nil
}

// SetNow Sets an element to our tskv, at the current time
func (t *Tskv) SetNow(value interface{}) error {
	t.Set(time.Now(), value)
	return nil
}
