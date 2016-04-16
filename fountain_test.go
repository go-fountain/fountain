// Package main implements a command-line utility for Fountain, a plain text markup language for screenwriting.
package main

import (
	"os"
	"testing"
)

func TestFountain(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd",
		"format", "testdata/hello-world.fountain",
	}
	main()
}

// TODO(ckaye) test exporting `testdata/hello-world.fountain` as text to terminal

func TestHelpCmd(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd",
		"-h",
	}
	main()
}
