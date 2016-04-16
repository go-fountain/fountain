// A screenplay element might be an Action, Character, Dialogue, etc.
package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "Generic", Generic.String())
	assert.Equal(t, "Title Page", TitlePage.String())
	assert.Equal(t, "Page Break", PageBreak.String())
	assert.Equal(t, "Action", Action.String())
	assert.Equal(t, "Centered Text", CenteredText.String())
	assert.Equal(t, "Character", Character.String())
	assert.Equal(t, "Dialogue", Dialogue.String())
	assert.Equal(t, "Dual Dialogue", DualDialogue.String())
	assert.Equal(t, "Emphasis", Emphasis.String())
	assert.Equal(t, "Lyrics", Lyrics.String())
	assert.Equal(t, "Parenthetical", Parenthetical.String())
	assert.Equal(t, "Scene Heading", SceneHeading.String())
	assert.Equal(t, "Transition", Transition.String())
	assert.Equal(t, "Comment", Comment.String())
	assert.Equal(t, "Note", Note.String())
	assert.Equal(t, "Section", Section.String())
	assert.Equal(t, "Synopsis", Synopsis.String())
	assert.Equal(t, "-", Class(12377877852).String())
}
