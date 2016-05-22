package tskv

// Element Smallest Atom
type Element struct {
	value interface{}
}

// Value returns integer form of Element
func (e *Element) Value() interface{} {
	return e.value
}

// Int returns integer form of Element
func (e *Element) Int() int {
	return e.value.(int)
}

/*
// String returns string form of Element
func (e *Element) String() string {
	return e.value.(string)
}
*/
