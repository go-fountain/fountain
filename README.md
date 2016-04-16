[![Build Status](https://travis-ci.org/go-fountain/fountain.svg?branch=master)](https://travis-ci.org/go-fountain/fountain) [![GoDoc](https://godoc.org/github.com/go-fountain/fountain?status.svg)](https://godoc.org/github.com/go-fountain/fountain) [![Go Report Card](https://goreportcard.com/badge/github.com/go-fountain/fountain)](https://goreportcard.com/report/github.com/go-fountain/fountain) [![codebeat badge](https://codebeat.co/badges/ac45ed94-c337-45bd-aabc-0fe952bae82c)](https://codebeat.co/projects/github-com-go-fountain-fountain)

[https://github.com/go-fountain/fountain](github.com/go-fountain/fountain)

## Fountain is a plain text markup language for screenwriting.

There's an example command-line utility `fountain.go` to demo the libraries, with a `bin/` wrapper.

To build and install `fountain` to your machine:

    make install

Then, to format a screenplay to terminal:

    $ fountain format hello-world.fountain
    
    ...................
    ...................
    ...................
    ...................
    ...................
    ...................

##### Credit

[Charney Kaye](http://w.charney.io)

[Outright Mental](http://w.outright.io)

## [Element](element/)

An Element is a single block of content, spanning the full width of the page, for a limited vertical area.

[![GoDoc](https://godoc.org/github.com/go-fountain/fountain/element?status.svg)](https://godoc.org/github.com/go-fountain/fountain/element) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-fountain/fountain/element)

## [Page](page/)

A Page is a vertical stack of Elements. Pagination is the ability to create a set of optimal pages from a set of elements.

[![GoDoc](https://godoc.org/github.com/go-fountain/fountain/page?status.svg)](https://godoc.org/github.com/go-fountain/fountain/page) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-fountain/fountain/page)

## [Export](export/)

An Export is a .PDF file created from a set of Pages paginated from a set of Elements.

[![GoDoc](https://godoc.org/github.com/go-fountain/fountain/export?status.svg)](https://godoc.org/github.com/go-fountain/fountain/export) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-fountain/fountain/export)
