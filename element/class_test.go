// A screenplay element might be an Action, Character, Dialogue, etc.
package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringOf(t *testing.T) {
	assert.Equal(t, "Generic", stringOf(Generic))
	assert.Equal(t, "Title Page", stringOf(TitlePage))
	assert.Equal(t, "Page Break", stringOf(PageBreak))
	assert.Equal(t, "Action", stringOf(Action))
	assert.Equal(t, "Centered Text", stringOf(CenteredText))
	assert.Equal(t, "Character", stringOf(Character))
	assert.Equal(t, "Dialogue", stringOf(Dialogue))
	assert.Equal(t, "Dual Dialogue", stringOf(DualDialogue))
	assert.Equal(t, "Emphasis", stringOf(Emphasis))
	assert.Equal(t, "Lyrics", stringOf(Lyrics))
	assert.Equal(t, "Parenthetical", stringOf(Parenthetical))
	assert.Equal(t, "Scene Heading", stringOf(SceneHeading))
	assert.Equal(t, "Transition", stringOf(Transition))
	assert.Equal(t, "Comment", stringOf(Comment))
	assert.Equal(t, "Note", stringOf(Note))
	assert.Equal(t, "Section", stringOf(Section))
	assert.Equal(t, "Synopsis", stringOf(Synopsis))
	assert.Equal(t, "-", stringOf(12377877852))
}
