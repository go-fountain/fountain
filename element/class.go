// A screenplay element might be an Action, Character, Dialogue, etc.
package element

// Class of pitch for a element (across all octaves)
type Class int

const (
	Generic Class = 0

	TitlePage Class = 101
	PageBreak Class = 102

	Action        Class = 201
	CenteredText  Class = 202
	Character     Class = 203
	Dialogue      Class = 204
	DualDialogue  Class = 205
	Emphasis      Class = 206
	Lyrics        Class = 207
	Parenthetical Class = 208
	SceneHeading  Class = 209
	Transition    Class = 210

	Comment  Class = 211
	Note     Class = 212
	Section  Class = 213
	Synopsis Class = 214
)

// String of the element, expressed with Sharps or Flats
func (from Class) String() string {
	return stringOf(from)
}

/*
 *
 private */

func stringOf(from Class) string {
	switch from {
	case Generic:
		return "Generic"
	case TitlePage:
		return "Title Page"
	case PageBreak:
		return "Page Break"
	case Action:
		return "Action"
	case CenteredText:
		return "Centered Text"
	case Character:
		return "Character"
	case Dialogue:
		return "Dialogue"
	case DualDialogue:
		return "Dual Dialogue"
	case Emphasis:
		return "Emphasis"
	case Lyrics:
		return "Lyrics"
	case Parenthetical:
		return "Parenthetical"
	case SceneHeading:
		return "Scene Heading"
	case Transition:
		return "Transition"
	case Comment:
		return "Comment"
	case Note:
		return "Note"
	case Section:
		return "Section"
	case Synopsis:
		return "Synopsis"
	}
	return "-"
}
