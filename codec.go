// Package qcodec provides encoding API definition and with several commonly
// used Codec such as uint32 and uint64 etc.
package qcodec

import (
	"errors"
	"reflect"
)

var (
	// ErrNotSlice indicates it expects a slice type but not
	ErrNotSlice = errors.New("it is not a slice")
	// ErrUnknownEltType indicates a type this package does not support.
	ErrUnknownEltType = errors.New("element type is unknown")

	// ErrNotFixedSize indicates the size of value of a type can not be
	// determined by its type.
	// Such slice of interface.
	ErrNotFixedSize = errors.New("element type is not fixed size")
)

// A Codec converts one element between serialized byte stream
// and in-memory data structure.
type Codec interface {
	// Convert into serialized byte stream.
	Encode(interface{}) []byte

	// Read byte stream and convert it back to typed data.
	Decode([]byte) (int, interface{})

	// Size returns the size in byte after encoding v.
	// If v is of type this encoder can not qcodec, it panics.
	Size(v interface{}) int

	// EncodedSize returns size of the encoded value.
	// Encoded element may be var-length.
	// This function is used to determine element size without the need of
	// encoding it.
	EncodedSize([]byte) int
}

// CodecOf returns a `Codec` implementation for type `e`
func CodecOf(e interface{}) (Codec, error) {
	k := reflect.ValueOf(e).Kind()
	return CodecByKind(k)
}

// GetSliceEltCodec creates a `Codec` for type of element in slice `s`
func GetSliceEltCodec(s interface{}) (Codec, error) {
	sl := reflect.ValueOf(s)
	if sl.Kind() != reflect.Slice {
		return nil, ErrNotSlice
	}

	eltKind := reflect.TypeOf(s).Elem().Kind()

	return CodecByKind(eltKind)
}

func CodecByKind(k reflect.Kind) (Codec, error) {
	var m Codec
	switch k {
	case reflect.Uint8:
		m = U8{}
	case reflect.Uint16:
		m = U16{}
	case reflect.Uint32:
		m = U32{}
	case reflect.Uint64:
		m = U64{}
	case reflect.Int8:
		m = I8{}
	case reflect.Int16:
		m = I16{}
	case reflect.Int32:
		m = I32{}
	case reflect.Int64:
		m = I64{}
	default:
		return nil, ErrUnknownEltType
	}

	return m, nil
}

// String16 converts uint16 to slice of 2 bytes and back.
type String16 struct{}

// Encode converts uint16 to slice of 2 bytes.
func (s String16) Encode(d interface{}) []byte {
	ss := d.(string)
	l := len(ss)
	rst := make([]byte, 2, 2+l)
	rst[0] = byte(l >> 8)
	rst[1] = byte(l)
	return append(rst, []byte(ss)...)
}

// Decode converts slice of 2 bytes to uint16.
// It returns number bytes consumed and an uint16.
func (s String16) Decode(b []byte) (int, interface{}) {
	l := int(b[0])<<8 + int(b[1])
	ss := string(b[2 : 2+l])
	return 2 + l, ss
}

// Size returns number of byte required to qcodec a string.
// It is len(str) + 2;
func (s String16) Size(d interface{}) int {
	ss := d.(string)
	l := len(ss)
	return 2 + l
}

// EncodedSize returned size of encoded data.
func (s String16) EncodedSize(b []byte) int {
	l := int(b[0])<<8 + int(b[1])
	return 2 + l
}
