package tskv

import "time"

// GetSince Returns all of the elements
func (t *Tskv) GetSince(when time.Time) *Tskv {
	toreturn, _ := New()
	for key, value := range t.elements {
		if key.After(when) {
			toreturn.Insert(key, value.Value())
		}
	}
	return toreturn
}
