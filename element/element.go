// An Element is a single block of content, spanning the full width of the page, for a limited vertical area.
//
// Credit
//
// Charney Kaye
// <hiya@charney.io>
// http://w.charney.io
//
// Outright Mental
// http://w.outright.io
//
package element

// Element models a musical element
type Element struct {
	Class  Class  // Class of pitch
}

// Elements is a set of elements, e.g. derived by parsing a screenplay file
type Elements []*Element

// OfClass pitch returns a Element model
func OfClass(class Class) (n *Element) {
	n = &Element{}
	n.Class = class
	return
}
