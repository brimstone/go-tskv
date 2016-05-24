package tskv

import "time"

// GetSince Returns all of the elements
func (t *Tskv) GetSince(when time.Time) *Tskv {
	t.cleanup()
	toreturn, _ := New(&Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	for key, value := range t.elements {
		if key.After(when) {
			toreturn.Insert(key, value.Value())
		}
	}
	return toreturn
}
