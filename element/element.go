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

// Element models an element of a screenplay
type Element struct {
	Text  string
	Class Class // Class of element
}

// Elements is a set of elements, e.g. derived by parsing a screenplay file
type Elements []*Element

// OfClass returns an Element model
func OfClass(class Class) (e *Element) {
	e = &Element{}
	e.Class = class
	return
}

// FromText parses text to return an Element model
func FromText(text string) (e *Element) {
	e = &Element{}
	e.Text = text
	return
}

/*
 *
 private */

func panicIf(e error) {
	if e != nil {
		panic(e)
	}
}
