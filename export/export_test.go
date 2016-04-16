// An Export is a .PDF file created from a set of Pages paginated from a set of Elements.
package export

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {
	p := &Export{}
	assert.NotNil(t, p)
}
