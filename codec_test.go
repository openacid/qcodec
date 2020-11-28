package qcodec

import (
	"testing"
)

func TestString16(t *testing.T) {

	cases := []struct {
		input string
		want  int
	}{
		{"", 2},
		{"a", 3},
		{"abc", 5},
	}

	m := String16{}

	for i, c := range cases {
		rst := m.Encode(c.input)
		if len(rst) != c.want {
			t.Fatalf("%d-th: encoded len: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, len(rst))
		}

		l := m.EncodedSize(rst)
		if l != c.want {
			t.Fatalf("%d-th: encoded size: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, l)
		}

		n, s := m.Decode(rst)
		if c.want != n {
			t.Fatalf("%d-th: decoded size: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, n)
		}
		if c.input != s {
			t.Fatalf("%d-th: decode: input: %v; want: %v; actual: %v",
				i+1, c.input, c.input, s)
		}
	}
}

func TestGetEncoder(t *testing.T) {

	cases := []struct {
		input   interface{}
		want    Codec
		wanterr error
	}{
		{ uint8(0), U8{}, nil, },
		{ uint16(0), U16{}, nil, },
		{ uint32(0), U32{}, nil, },
		{ uint64(0), U64{}, nil, },
		{ int16(0), I16{}, nil, },
		{ int8(0), I8{}, nil, },
		{ int32(0), I32{}, nil, },
		{ int64(0), I64{}, nil, },
		{ []int{}, nil, ErrUnknownEltType, },
		{ nil, nil, ErrUnknownEltType, },
	}

	for i, c := range cases {
		rst, err := CodecOf(c.input)
		if rst != c.want {
			t.Fatalf("%d-th: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, rst)
		}
		if err != c.wanterr {
			t.Fatalf("%d-th: input: %v; wantErr: %v; actual: %v",
				i+1, c.input, c.wanterr, err)
		}
	}
}

func TestGetSliceEltEncoder(t *testing.T) {

	cases := []struct {
		input   interface{}
		want    Codec
		wanterr error
	}{
		{
			[]uint16{},
			U16{},
			nil,
		},
		{
			[]uint32{},
			U32{},
			nil,
		},
		{
			[]uint64{},
			U64{},
			nil,
		},
		{
			[]int{},
			nil,
			ErrUnknownEltType,
		},
		{
			int(1),
			nil,
			ErrNotSlice,
		},
	}

	for i, c := range cases {
		rst, err := GetSliceEltCodec(c.input)
		if rst != c.want {
			t.Fatalf("%d-th: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, rst)
		}
		if err != c.wanterr {
			t.Fatalf("%d-th: input: %v; wantErr: %v; actual: %v",
				i+1, c.input, c.wanterr, err)
		}
	}
}
