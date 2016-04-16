// Fountain is a plain text markup language for screenwriting.
//
// There's an example command-line utility `fountain.go` to demo the libraries, with a `bin/` wrapper.
//
// Build and install `fountain` to your machine
//
//     make install
//
// Credit
//
// Charney Kaye
// <hiya@charney.io>
// http://w.charney.io
//
// Outright Mental
// http://w.outright.io
//
package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "fountain"
	app.Usage = "a plain text markup language for screenwriting"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{cli.Author{Name: "Charney Kaye", Email: "hiya@charney.io"}}
	app.Commands = commands
	app.Run(os.Args)
}

var commands = []cli.Command{

	{ // Format paginated script to terminal
		Name:        "format",
		Aliases:     []string{"fmt","f"},
		Usage:       "format a screenplay to terminal",
		Description: "A formatted screenplay is a set of Pages exported as a raster of characters.",
		Action: func(c *cli.Context) {
			file := c.Args().First()
			if len(file) > 0 {
				// TODO(charneykaye) output the formatted screenplay to terminal
			} else {
				// no arguments
				cli.ShowCommandHelp(c, "format")
			}
		},
	},

}
