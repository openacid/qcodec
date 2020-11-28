package qcodec

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"github.com/pkg/errors"
)

// defaultEndian is default endian
var defaultEndian = binary.LittleEndian

// TypeCodec provides encoding for fixed size types.
// Such as int32 or struct { X int32; Y int64; }
//
// "int" is not a fixed size type: int on different platform has different size,
// 4 or 8 bytes.
//
// "[]int32" is not a fixed size type: the data size is also defined by the
// number of elements.
type TypeCodec struct {
	// byteOrder defines the byte order to qcodec a value.
	// By default it is binary.LittleEndian
	byteOrder binary.ByteOrder
	// typ is the data type to qcodec.
	typ reflect.Type
	// size is the encoded size of this type.
	size int
}

// NewTypeCodec creates a *TypeCodec by a value.
// The value "zero" defines what type this Codec can deal with and must be a
// fixed size type.
// "endian" could be binary.LittleEndian or binary.BigEndian.
func NewTypeCodec(zero interface{}, endians ...binary.ByteOrder) (*TypeCodec, error) {
	var endian binary.ByteOrder = nil
	if len(endians) > 0 {
		endian = endians[0]
	}

	if endian == nil {
		endian = defaultEndian
	}

	m := &TypeCodec{
		byteOrder: endian,
		typ:       reflect.Indirect(reflect.ValueOf(zero)).Type(),
		size:      binary.Size(zero),
	}

	if m.size == -1 {
		return nil, errors.Wrapf(ErrNotFixedSize, "type: %v", reflect.TypeOf(zero))
	}
	if m.typ.Kind() == reflect.Slice {
		return nil, errors.Wrapf(ErrNotFixedSize, "slice size is not fixed")
	}

	return m, nil
}

// NewTypeCodecByType creates a *TypeCodec for specified type and with a specified byte order.
//
// "endian" could be binary.LittleEndian or binary.BigEndian.
func NewTypeCodecByType(t reflect.Type, endian binary.ByteOrder) (*TypeCodec, error) {
	v := reflect.New(t)
	return NewTypeCodec(v.Interface(), endian)
}

// Encode converts a m.typ value to byte slice.
// If a different type value from the one used with NewTypeCodec passed in,
// it panics.
func (m *TypeCodec) Encode(d interface{}) []byte {
	if reflect.Indirect(reflect.ValueOf(d)).Type() != m.typ {
		panic("different type from TypeCodec.typ")
	}

	b := bytes.NewBuffer(make([]byte, 0, m.size))
	err := binary.Write(b, m.byteOrder, d)
	if err != nil {
		// there should not be any error if type is fixed size
		panic(err)
	}
	return b.Bytes()
}

// Decode converts byte slice to a pointer to typ value.
// It returns number bytes consumed and an typ value in interface{}.
func (m *TypeCodec) Decode(b []byte) (int, interface{}) {

	b = b[0:m.size]
	v := reflect.New(m.typ)
	err := binary.Read(bytes.NewBuffer(b), m.byteOrder, v.Interface())
	if err != nil {
		panic(err)
	}
	return m.size, reflect.Indirect(v).Interface()
}

// GetSize returns m.size.
func (m *TypeCodec) Size(d interface{}) int {
	return m.size
}

// GetEncodedSize returns m.size.
func (m *TypeCodec) EncodedSize(b []byte) int {
	return m.size
}
