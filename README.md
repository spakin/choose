choose
======

[![Build Status](https://travis-ci.org/spakin/choose.svg?branch=master)](https://travis-ci.org/spakin/choose)
[![Go Report Card](https://goreportcard.com/badge/github.com/spakin/choose)](https://goreportcard.com/report/github.com/spakin/choose)
[![GoDoc](https://godoc.org/github.com/spakin/choose?status.svg)](https://godoc.org/github.com/spakin/choose)

Introduction
------------

`choose` is a small package for the [Go Programming Language](https://golang.org/) that produces all combinations of *M* out of *N* items.  The core part of the code is based on [Matthew Belmonte's C implementation](http://www.netlib.org/toms-2014-06-10/382) of Phillip J. Chase's [Algorithm 382: Combinations of *M* out of *N* Objects](https://doi.org/10.1145/362384.362502).

Installation
------------

The `choose` package has opted into the [Go module system](https://blog.golang.org/using-go-modules) so installation is in fact unnecessary if your program or package has done likewise.  Otherwise, a traditional
```bash
go get github.com/spakin/choose
```
will install the package.

Documentation
-------------

See the [`choose` API reference](https://pkg.go.dev/github.com/spakin/choose?tab=doc) for details, but the basic usage model is to invoke a function that accepts a slice and returns a channel, then iterate over that channel's contents, each of which is a slice of the same type as the input.

Author
------

[Scott Pakin](https://www.pakin.org/~scott/), [*scott+choose@pakin.org*](mailto:scott+choose@pakin.org)
