package tskv

import "time"

// Cleanup
func (t *Tskv) cleanup() error {
	// add item to list
	// if duration is shorter than list
	// remove last item
	t.youngest = time.Now().Truncate(t.frequency)
	t.oldest = t.youngest.Add(-t.duration)
	for k := range t.elements {
		if k.Before(t.oldest) {
			delete(t.elements, k)
		}
	}
	return nil
}
