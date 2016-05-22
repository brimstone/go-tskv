package tskv

import "time"

// Insert Inserts an element to our tskv
func (t *Tskv) Insert(when time.Time, value interface{}) error {
	// add item to list
	// if duration is shorter than list
	// remove last item
	t.elements[when] = Element{value: value}
	return nil
}

// InsertNow Inserts an element to our tskv, at the current time
func (t *Tskv) InsertNow(value interface{}) error {
	// add item to list
	// if duration is shorter than list
	// remove last item
	t.elements[time.Now()] = Element{value: value}
	return nil
}
