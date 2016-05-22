package tskv

import "time"

// Oldest returns the oldest element in our tskv
func (t *Tskv) Oldest() *Element {
	oldestAge := time.Now()
	var oldestElement Element

	for age, value := range t.elements {
		if age.Before(oldestAge) {
			oldestAge = age
			oldestElement = value
		}
	}
	return &oldestElement
}
