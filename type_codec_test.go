package qcodec

import (
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

type typeXY struct {
	X int32
	Y int32
}

func testPanic(t *testing.T, f func(), msg string) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic: %s", msg)
		}
	}()

	f()
}

func TestNewTypeEncoder(t *testing.T) {

	m, _ := NewTypeCodec(int32(1))
	if m.byteOrder != binary.LittleEndian {
		t.Fatalf("expect default endian is %#v but %#v", binary.LittleEndian, m.byteOrder)
	}

	ii := int32(1)

	cases := []struct {
		input   interface{}
		want    *TypeCodec
		wantErr error
	}{
		{
			1,
			nil,
			ErrNotFixedSize,
		},
		{
			[]int32{1},
			nil,
			ErrNotFixedSize,
		},
		{
			int32(1),
			&TypeCodec{
				byteOrder: binary.LittleEndian,
				typ:       reflect.ValueOf(int32(1)).Type(),
				size:      4,
			},
			nil,
		},
		{
			&ii,
			&TypeCodec{
				byteOrder: binary.LittleEndian,
				typ:       reflect.ValueOf(int32(1)).Type(),
				size:      4,
			},
			nil,
		},
		{
			typeXY{1, 2},
			&TypeCodec{
				byteOrder: binary.LittleEndian,
				typ:       reflect.ValueOf(typeXY{}).Type(),
				size:      8,
			},
			nil,
		},
		{
			&typeXY{1, 2},
			&TypeCodec{
				byteOrder: binary.LittleEndian,
				typ:       reflect.ValueOf(typeXY{}).Type(),
				size:      8,
			},
			nil,
		},
	}

	for i, c := range cases {
		rst, err := NewTypeCodec(c.input, nil)
		if errors.Cause(err) != c.wantErr {
			t.Fatalf("%d-th: input: %#v; wantErr: %#v; actual: %#v",
				i+1, c.input, c.wantErr, err)
		}

		if !reflect.DeepEqual(c.want, rst) {
			t.Fatalf("%d-th: input: %#v; want: %#v; actual: %#v",
				i+1, c.input, c.want, rst)
		}

		m, err := NewTypeCodecByType(
			reflect.Indirect(reflect.ValueOf(c.input)).Type(), nil)
		if errors.Cause(err) != c.wantErr {
			t.Fatalf("%d-th: input: %#v; wantErr: %#v; actual: %#v",
				i+1, c.input, c.wantErr, err)
		}

		if !reflect.DeepEqual(c.want, m) {
			t.Fatalf("%d-th: input: %#v; want: %#v; actual: %#v",
				i+1, c.input, c.want, m)
		}
	}
}

func TestTypeEncoderEncode(t *testing.T) {

	m, err := NewTypeCodec(int32(1), nil)
	if err != nil {
		t.Fatalf("expected no error but: %v", err)
	}

	testPanic(t, func() { m.Encode(uint32(1)) }, "int32: uint32")
	testPanic(t, func() { m.Encode([]int32{1}) }, "int32: []int32")

	// indirect value results in no panic
	ii := int32(5)
	bs := m.Encode(&ii)
	want := []byte{5, 0, 0, 0}
	if !reflect.DeepEqual(want, bs) {
		t.Fatalf("want: %#v, but: %#v", want, bs)
	}

	cases := []struct {
		input interface{}
		want  []byte
	}{
		{
			int32(1),
			[]byte{1, 0, 0, 0},
		},
		{
			byte(1),
			[]byte{1},
		},
		{
			typeXY{1, 2},
			[]byte{1, 0, 0, 0, 2, 0, 0, 0},
		},
		{
			&typeXY{1, 2},
			[]byte{1, 0, 0, 0, 2, 0, 0, 0},
		},
	}

	for i, c := range cases {
		m, err := NewTypeCodec(c.input, nil)
		if err != nil {
			t.Fatalf("%d-th: expected no error but: %#v", i+1, err)
		}

		n := m.Size(c.input)
		if n != binary.Size(c.input) {
			t.Fatalf("expect n to be %d but %d", binary.Size(c.input), n)
		}

		bs := m.Encode(c.input)
		if !reflect.DeepEqual(c.want, bs) {
			t.Fatalf("%d-th: input: %#v; want: %#v; actual: %#v",
				i+1, c.input, c.want, bs)
		}
	}
}

func TestTypeEncoderDecode(t *testing.T) {

	ii := int32(5)

	cases := []struct {
		input interface{}
		want  interface{}
	}{
		{
			int32(1),
			int32(1),
		},
		{
			byte(1),
			byte(1),
		},
		{
			&ii,
			int32(5),
		},
		{
			typeXY{1, 2},
			typeXY{1, 2},
		},
		{
			&typeXY{1, 2},
			typeXY{1, 2},
		},
	}

	for i, c := range cases {
		m, err := NewTypeCodec(c.input, nil)
		if err != nil {
			t.Fatalf("%d-th: expected no error but: %#v", i+1, err)
		}

		bs := m.Encode(c.input)
		n, v := m.Decode(bs)

		if n != m.size {
			t.Fatalf("expect n to b %d but %d", m.size, n)
		}

		if n != m.EncodedSize(bs) {
			t.Fatalf("expect n to b %d but %d", m.EncodedSize(bs), n)
		}

		if !reflect.DeepEqual(c.want, v) {
			t.Fatalf("%d-th: input: %#v; want: %#v; actual: %#v",
				i+1, c.input, c.want, v)
		}
	}
}
