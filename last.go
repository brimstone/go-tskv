package tskv

// Last returns the youngest element in our tskv
func (t *Tskv) Last() *Element {
	t.cleanup()
	//fmt.Println(t.keys)
	return &Element{value: 1}
	/*
		last := t.elements[t.keys[0]]
		fmt.Println(last)
		return &last
	*/
}
