# trie
A package for trie tree implementations.

[![GoDoc](https://godoc.org/github.com/drexedam/trie?status.svg)](https://godoc.org/github.com/drexedam/trie)
[![Go Report Card](https://goreportcard.com/badge/github.com/drexedam/trie)](https://goreportcard.com/report/github.com/drexedam/trie)
[![Build Status](https://travis-ci.org/drexedam/trie.svg?branch=master)](https://travis-ci.org/drexedam/trie)

## Install
```sh
go get -u github.com/drexedam/trie
```

## Available Implementations
Currently only a straight forward implementation using maps for storing children is available.

## Examples
```go
tr := New()
tr.Insert("Test")
if tr.Find("Test") {
        // Found
}
```

## License
MIT licensed. See the LICENSE file for details.