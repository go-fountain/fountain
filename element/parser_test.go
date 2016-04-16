// Parse a .fountain file into screenplay Elements
package element

import (
	"io/ioutil"
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

// TODO(charneykaye) test parse file into elements

func TestParseFile(t *testing.T) {
	RegisterTestingT(t)
	expectations := getExpectationsOf("testdata/hello-world.expectations.yaml")
	elements := ParseFile("testdata/hello-world.fountain")
	Expect(elements).To(ConsistOf(expectations.Elements))
}

func TestParser_NextByte(t *testing.T) {
	// TODO(ckaye) test parsing next byte
}

type TestParserExpectations struct {
	Elements Elements `yaml:"elements"`
}

/*
 *
 */

func getExpectationsOf(filename string) TestParserExpectations {
	expectations := TestParserExpectations{}
	data, err := ioutil.ReadFile(filename)
	Expect(err).To(BeNil())
	err = yaml.Unmarshal(data, &expectations)
	Expect(err).To(BeNil())
	return expectations
}
