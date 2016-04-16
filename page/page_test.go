// A Page is a vertical stack of Elements. Pagination is the ability to create a set of optimal pages from a set of elements.
package page

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	p := &Page{}
	assert.NotNil(t, p)
}
