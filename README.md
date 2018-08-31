<h1 align="center">
    <img src="https://github.com/gurre/detectreader/blob/master/gopher_detectreader.png" alt="Mascot" width="300">
    <br />
    Detect Reader
</h1>

[![GoDoc](https://godoc.org/github.com/gurre/detectreader?status.svg)](https://godoc.org/github.com/gurre/detectreader)
[![License](http://img.shields.io/:license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/gurre/detectreader)](https://goreportcard.com/report/github.com/gurre/detectreader)

Detect reader reads the first bytes looking for magic bytes. Then it returns an appropriate reader for that stream.

## Installation

Windows, OS X & Linux:

```
go get github.com/gurre/detectreader
```

## Usage examples

Simple usage:

```go
// Some compressed byte stream
reader := bytes.NewReader(b)
// The returned reader is a decompressor
decompressedReader, err := Decompress(reader)
```
