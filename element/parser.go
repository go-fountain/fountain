// Parse a .fountain file into screenplay Elements

package element

import (
	"io/ioutil"
)

func ParseFile(filename string) (elements Elements) {
	// TODO(charneykaye) parse file into elements

	data, err := ioutil.ReadFile(filename)
	panicIf(err)

	parse := &Parser{}
	for _, b := range data {
		parse.NextByte(b)
	}
	elements = parse.Elements

	return
}

type Parser struct {
	Elements Elements

	elementInProgress []byte
}

func (self *Parser) NextByte(b byte) {
	switch b {
	case 0x0A, 0x0D:
		self.finishElementInProgress()
	default:
		self.elementInProgress = append(self.elementInProgress, b)
	}
}

/*
 *
 private */

func (self *Parser) finishElementInProgress() {
	text := string(self.elementInProgress[:])
	if len(text) > 0 {
		self.Elements = append(self.Elements, FromText(text))
	}
	self.elementInProgress = make([]byte, 0)
}
