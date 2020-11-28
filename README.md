# qcodec

[![Travis](https://travis-ci.com/openacid/qcodec.svg?branch=main)](https://travis-ci.com/openacid/qcodec)
![test](https://github.com/openacid/qcodec/workflows/test/badge.svg)

[![Report card](https://goreportcard.com/badge/github.com/openacid/qcodec)](https://goreportcard.com/report/github.com/openacid/qcodec)
[![Coverage Status](https://coveralls.io/repos/github/openacid/qcodec/badge.svg?branch=main&service=github)](https://coveralls.io/github/openacid/qcodec?branch=main&service=github)

[![GoDoc](https://godoc.org/github.com/openacid/qcodec?status.svg)](http://godoc.org/github.com/openacid/qcodec)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/openacid/qcodec)](https://pkg.go.dev/github.com/openacid/qcodec)
[![Sourcegraph](https://sourcegraph.com/github.com/openacid/qcodec/-/badge.svg)](https://sourcegraph.com/github.com/openacid/qcodec?badge)

qcodec encodes and decodes a primitive type or struct type with a predefined codec engine.
It is simple thus very fast.


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Install](#install)
- [Synopsis](#synopsis)
  - [Build a SlimArray](#build-a-slimarray)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Install

```sh
go get github.com/openacid/qcodec
```

# Synopsis

## Build a SlimArray

```go
package qcodec

import "fmt"

func Example() {
	type Foo struct {
		A int32
		B int64
	}

	fmt.Println("uint32:", U32{}.Encode(uint32(0x0102)))
	fmt.Println("uint16:", U16{}.Encode(uint16(0x0102)))
	fmt.Println("dummy:", Dummy{}.Encode("dummy encodes anything to empty slice"))

	// TypeCodec encodes any size fixed type.
	fcodec, err := NewTypeCodec(Foo{})
	_ = err
	buf := fcodec.Encode(Foo{1, 2})
	fmt.Println("Foo:", buf)
	consumed, foo := fcodec.Decode(buf)
	fmt.Println("decode Foo:", consumed, foo)

	intCodec, err := NewTypeCodec(int(1))
	_ = intCodec
	fmt.Println("int:", err)

	// Output:
	// uint32: [2 1 0 0]
	// uint16: [2 1]
	// dummy: []
	// Foo: [1 0 0 0 2 0 0 0 0 0 0 0]
	// decode Foo: 12 {1 2}
	// int: type: int: element type is not fixed size
}
```