package main

import (
	"github.com/openacid/genr"
)

var implHead = `package qcodec

import "encoding/binary"
`

var implTemplate = `
// {{.Name}} converts {{.ValType}} to slice of {{.ValLen}} bytes and back.
type {{.Name}} struct{}

// Encode converts {{.ValType}} to slice of {{.ValLen}} bytes.
func (c {{.Name}}) Encode(d interface{}) []byte {
	b := make([]byte, {{.ValLen}})
	v := {{.EncodeCast}}(d.({{.ValType}}))
	binary.LittleEndian.Put{{.Codec}}(b, v)
	return b
}

// Decode converts slice of {{.ValLen}} bytes to {{.ValType}}.
// It returns number bytes consumed and an {{.ValType}}.
func (c {{.Name}}) Decode(b []byte) (int, interface{}) {

	size := int({{.ValLen}})
	s := b[:size]

	d := {{.ValType}}(binary.LittleEndian.{{.Codec}}(s))
	return size, d
}

// Size returns the size in byte after encoding v.
func (c {{.Name}}) Size(d interface{}) int {
	return {{.ValLen}}
}

// EncodedSize returns {{.ValLen}}.
func (c {{.Name}}) EncodedSize(b []byte) int {
	return {{.ValLen}}
}
`

var testHead = `package qcodec

import (
	"testing"
)
`

var testTemplate = `
func Test{{.Name}}(t *testing.T) {

	v0 := [8]byte{}
	v1 := [8]byte{1}
	v1234 := [8]byte{0x34, 0x12}
	vneg := [8]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	cases := []struct {
		input    {{.ValType}}
		want     string
		wantsize int
	}{
		{0, string(v0[:{{.ValLen}}]), {{.ValLen}}},
		{1, string(v1[:{{.ValLen}}]), {{.ValLen}}},
		{0x1234, string(v1234[:{{.ValLen}}]), {{.ValLen}}},
		{^{{.ValType}}(0), string(vneg[:{{.ValLen}}]), {{.ValLen}}},
	}

	m := {{.Name}}{}

	for i, c := range cases {
		rst := m.Encode(c.input)
		if string(rst) != c.want {
			t.Fatalf("%d-th: input: %v; want: %v; actual: %v",
				i+1, c.input, []byte(c.want), rst)
		}

		n := m.Size(c.input)
		if c.wantsize != n {
			t.Fatalf("%d-th: input: %v; wantsize: %v; actual: %v",
				i+1, c.input, c.wantsize, n)
		}

		n = m.EncodedSize(rst)
		if c.wantsize != n {
			t.Fatalf("%d-th: input: %v; wantsize: %v; actual: %v",
				i+1, c.input, c.wantsize, n)
		}

		n, u64 := m.Decode(rst)
		if c.input != u64 {
			t.Fatalf("%d-th: decode: input: %v; want: %v; actual: %v",
				i+1, c.input, c.input, u64)
		}
		if c.wantsize != n {
			t.Fatalf("%d-th: decoded size: input: %v; want: %v; actual: %v",
				i+1, c.input, c.wantsize, n)
		}
	}
}
`

func main() {

	pref := "int"
	implfn := pref + ".go"
	testfn := pref + "_test.go"

	impls := []interface{}{
		genr.NewIntConfig("U16", "uint16"),
		genr.NewIntConfig("U32", "uint32"),
		genr.NewIntConfig("U64", "uint64"),
		genr.NewIntConfig("I16", "int16"),
		genr.NewIntConfig("I32", "int32"),
		genr.NewIntConfig("I64", "int64"),
	}

	genr.Render(implfn, implHead, implTemplate, impls, []string{"gofmt", "unconvert"})
	genr.Render(testfn, testHead, testTemplate, impls, []string{"gofmt", "unconvert"})
}
