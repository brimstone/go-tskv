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

// GetRaw returns an raw repssentation of the series
func (t *Tskv) GetRaw() []Element {
	t.cleanup()
	var all []Element
	current := time.Now().Truncate(t.frequency)
	oldest := current.Add(-t.duration)
	for current.After(oldest) {
		all = append(all, t.elements[current])
		current = current.Add(-t.frequency)
	}
	return all
}

// GetInt returns and integer representation of the series
func (t *Tskv) GetInt() []int {
	var all []int
	for _, raw := range t.GetRaw() {
		if raw.Value() == nil {
			all = append(all, 0)
		} else {
			all = append(all, raw.Int())
		}

	}
	return all
}
