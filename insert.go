package tskv

import "time"

// Insert Inserts an element to our tskv
func (t *Tskv) Insert(when time.Time, value interface{}) error {
	// add item to list
	// if duration is shorter than list
	// remove last item
	t.elements[when.Format(time.RFC3339)] = value
	return nil
}
