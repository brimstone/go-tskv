package tskv

import "time"

// Insert Inserts an element to our tskv
func (t *Tskv) Insert(when time.Time, value interface{}) error {
	t.cleanup()
	t.elements[when.Truncate(t.frequency)] = Element{value: value}
	return nil
}

// InsertNow Inserts an element to our tskv, at the current time
func (t *Tskv) InsertNow(value interface{}) error {
	t.Insert(time.Now(), value)
	return nil
}
