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

func TestHelpCmd(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd",
		"-h",
	}
	main()
}
