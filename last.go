package tskv

// Last returns the youngest element in our tskv
func (t *Tskv) Last() *Element {
	t.cleanup()
	last := t.elements[t.youngest]
	return &last
}
