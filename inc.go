package tskv

import "time"

// Inc increments an element to our tskv
func (t *Tskv) Inc(when time.Time, value interface{}) error {
	t.cleanup()
	// Get our current value
	current := t.elements[when.Truncate(t.frequency)]
	// If we don't have a value
	if current.value == nil {
		// Set it to a value
		t.Set(when, value)
		return nil
	}
	// update our element
	t.elements[when.Truncate(t.frequency)] = Element{
		value: current.value.(int) + value.(int),
	}
	return nil
}

// IncNow increments an element to our tskv, at the current time
func (t *Tskv) IncNow(value interface{}) error {
	t.Inc(time.Now(), value)
	return nil
}
