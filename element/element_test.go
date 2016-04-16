// An Element is a single block of content, spanning the full width of the page, for a limited vertical area.
package element

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

func TestOfClass(t *testing.T) {
	RegisterTestingT(t)
	n := OfClass(Action)
	Expect(n).To(Equal(&Element{
		Class: Action,
	}))
}

func TestPanicIf(t *testing.T) {
	RegisterTestingT(t)
	defer func(){
		r := recover()
		Expect(r).NotTo(BeNil())
	}()
	err := errors.New("test fail")
	panicIf(err)
}
