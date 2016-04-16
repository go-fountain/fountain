// An Element is a single block of content, spanning the full width of the page, for a limited vertical area.
package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfClass(t *testing.T) {
	n := OfClass(Action)
	assert.Equal(t, n, &Element{
		Class: Action,
	})
}
