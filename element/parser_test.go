// Parse a .fountain file into screenplay Elements
package element

import (
	"testing"

	. "github.com/onsi/gomega"
)

// TODO(charneykaye) test parse file into elements

func TestParseFile(t *testing.T) {
	RegisterTestingT(t)
	elements := ParseFile("testdata/hello-world.fountain")
	Expect(elements).To(HaveLen(29))
}

func TestParser_NextByte(t *testing.T) {
	// TODO(ckaye) test parsing next byte
}
