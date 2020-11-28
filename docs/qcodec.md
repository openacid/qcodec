# qcodec
--
    import "github.com/openacid/qcodec"

Package qcodec provides encoding API definition and with several commonly used
Codec such as uint32 and uint64 etc.

## Usage

```go
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
```

#### type Bytes

```go
type Bytes struct {
}
```

Bytes converts a byte slice into fixed length slice. Result slice length is
defined by Bytes.size .

#### func (Bytes) Decode

```go
func (c Bytes) Decode(b []byte) (int, interface{})
```
Decode copies fixed length slice out of source byte slice. The returned bytes
are NOT copied.

#### func (Bytes) Encode

```go
func (c Bytes) Encode(d interface{}) []byte
```
Encode converts byte slice to byte slice.

#### func (Bytes) EncodedSize

```go
func (c Bytes) EncodedSize(b []byte) int
```
GetEncodedSize returns c.size

#### func (Bytes) Size

```go
func (c Bytes) Size(d interface{}) int
```
GetSize returns the length: c.size.

#### type Codec

```go
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
```

A Codec converts one element between serialized byte stream and in-memory data
structure.

#### func  CodecByKind

```go
func CodecByKind(k reflect.Kind) (Codec, error)
```

#### func  CodecOf

```go
func CodecOf(e interface{}) (Codec, error)
```
CodecOf returns a `Codec` implementation for type `e`

#### func  GetSliceEltCodec

```go
func GetSliceEltCodec(s interface{}) (Codec, error)
```
GetSliceEltCodec creates a `Codec` for type of element in slice `s`

#### type Dummy

```go
type Dummy struct {
}
```

Dummy converts anything to nothing.

#### func (Dummy) Decode

```go
func (c Dummy) Decode(b []byte) (int, interface{})
```
Decode always returns nil.

#### func (Dummy) Encode

```go
func (c Dummy) Encode(d interface{}) []byte
```
Encode converts something to empty byte slice.

#### func (Dummy) EncodedSize

```go
func (c Dummy) EncodedSize(b []byte) int
```
GetEncodedSize returns 0.

#### func (Dummy) Size

```go
func (c Dummy) Size(d interface{}) int
```
GetSize returns 0.

#### type I16

```go
type I16 struct{}
```

I16 converts int16 to slice of 2 bytes and back.

#### func (I16) Decode

```go
func (c I16) Decode(b []byte) (int, interface{})
```
Decode converts slice of 2 bytes to int16. It returns number bytes consumed and
an int16.

#### func (I16) Encode

```go
func (c I16) Encode(d interface{}) []byte
```
Encode converts int16 to slice of 2 bytes.

#### func (I16) EncodedSize

```go
func (c I16) EncodedSize(b []byte) int
```
GetEncodedSize returns 2.

#### func (I16) Size

```go
func (c I16) Size(d interface{}) int
```
GetSize returns the size in byte after encoding v.

#### type I32

```go
type I32 struct{}
```

I32 converts int32 to slice of 4 bytes and back.

#### func (I32) Decode

```go
func (c I32) Decode(b []byte) (int, interface{})
```
Decode converts slice of 4 bytes to int32. It returns number bytes consumed and
an int32.

#### func (I32) Encode

```go
func (c I32) Encode(d interface{}) []byte
```
Encode converts int32 to slice of 4 bytes.

#### func (I32) EncodedSize

```go
func (c I32) EncodedSize(b []byte) int
```
GetEncodedSize returns 4.

#### func (I32) Size

```go
func (c I32) Size(d interface{}) int
```
GetSize returns the size in byte after encoding v.

#### type I64

```go
type I64 struct{}
```

I64 converts int64 to slice of 8 bytes and back.

#### func (I64) Decode

```go
func (c I64) Decode(b []byte) (int, interface{})
```
Decode converts slice of 8 bytes to int64. It returns number bytes consumed and
an int64.

#### func (I64) Encode

```go
func (c I64) Encode(d interface{}) []byte
```
Encode converts int64 to slice of 8 bytes.

#### func (I64) EncodedSize

```go
func (c I64) EncodedSize(b []byte) int
```
GetEncodedSize returns 8.

#### func (I64) Size

```go
func (c I64) Size(d interface{}) int
```
GetSize returns the size in byte after encoding v.

#### type I8

```go
type I8 struct{}
```

I8 converts int8 to slice of 1 byte and back.

#### func (I8) Decode

```go
func (c I8) Decode(b []byte) (int, interface{})
```
Decode converts slice of 1 byte to int8. It returns number bytes consumed and an
int8.

#### func (I8) Encode

```go
func (c I8) Encode(d interface{}) []byte
```
Encode converts int8 to slice of 1 byte.

#### func (I8) EncodedSize

```go
func (c I8) EncodedSize(b []byte) int
```
GetEncodedSize returns 2.

#### func (I8) Size

```go
func (c I8) Size(d interface{}) int
```
GetSize returns the size in byte after encoding v.

#### type Int

```go
type Int struct{}
```

Int converts int to slice of bytes and back.

#### func (Int) Decode

```go
func (c Int) Decode(b []byte) (int, interface{})
```
Decode converts slice of bytes to int. It returns number bytes consumed and an
int.

#### func (Int) Encode

```go
func (c Int) Encode(d interface{}) []byte
```
Encode converts int to slice of bytes.

#### func (Int) EncodedSize

```go
func (c Int) EncodedSize(b []byte) int
```
GetEncodedSize returns native int size.

#### func (Int) Size

```go
func (c Int) Size(d interface{}) int
```
GetSize returns native int size in byte after encoding v.

#### type String16

```go
type String16 struct{}
```

String16 converts uint16 to slice of 2 bytes and back.

#### func (String16) Decode

```go
func (s String16) Decode(b []byte) (int, interface{})
```
Decode converts slice of 2 bytes to uint16. It returns number bytes consumed and
an uint16.

#### func (String16) Encode

```go
func (s String16) Encode(d interface{}) []byte
```
Encode converts uint16 to slice of 2 bytes.

#### func (String16) EncodedSize

```go
func (s String16) EncodedSize(b []byte) int
```
EncodedSize returned size of encoded data.

#### func (String16) Size

```go
func (s String16) Size(d interface{}) int
```
Size returns number of byte required to qcodec a string. It is len(str) + 2;

#### type TypeCodec

```go
type TypeCodec struct {
}
```

TypeCodec provides encoding for fixed size types. Such as int32 or struct { X
int32; Y int64; }

"int" is not a fixed size type: int on different platform has different size, 4
or 8 bytes.

"[]int32" is not a fixed size type: the data size is also defined by the number
of elements.

#### func  NewTypeCodec

```go
func NewTypeCodec(zero interface{}, endians ...binary.ByteOrder) (*TypeCodec, error)
```
NewTypeCodec creates a *TypeCodec by a value. The value "zero" defines what type
this Codec can deal with and must be a fixed size type. "endian" could be
binary.LittleEndian or binary.BigEndian.

#### func  NewTypeCodecByType

```go
func NewTypeCodecByType(t reflect.Type, endian binary.ByteOrder) (*TypeCodec, error)
```
NewTypeCodecByType creates a *TypeCodec for specified type and with a specified
byte order.

"endian" could be binary.LittleEndian or binary.BigEndian.

#### func (*TypeCodec) Decode

```go
func (m *TypeCodec) Decode(b []byte) (int, interface{})
```
Decode converts byte slice to a pointer to typ value. It returns number bytes
consumed and an typ value in interface{}.

#### func (*TypeCodec) Encode

```go
func (m *TypeCodec) Encode(d interface{}) []byte
```
Encode converts a m.typ value to byte slice. If a different type value from the
one used with NewTypeCodec passed in, it panics.

#### func (*TypeCodec) EncodedSize

```go
func (m *TypeCodec) EncodedSize(b []byte) int
```
GetEncodedSize returns m.size.

#### func (*TypeCodec) Size

```go
func (m *TypeCodec) Size(d interface{}) int
```
GetSize returns m.size.

#### type U16

```go
type U16 struct{}
```

U16 converts uint16 to slice of 2 bytes and back.

#### func (U16) Decode

```go
func (c U16) Decode(b []byte) (int, interface{})
```
Decode converts slice of 2 bytes to uint16. It returns number bytes consumed and
an uint16.

#### func (U16) Encode

```go
func (c U16) Encode(d interface{}) []byte
```
Encode converts uint16 to slice of 2 bytes.

#### func (U16) EncodedSize

```go
func (c U16) EncodedSize(b []byte) int
```
GetEncodedSize returns 2.

#### func (U16) Size

```go
func (c U16) Size(d interface{}) int
```
GetSize returns the size in byte after encoding v.

#### type U32

```go
type U32 struct{}
```

U32 converts uint32 to slice of 4 bytes and back.

#### func (U32) Decode

```go
func (c U32) Decode(b []byte) (int, interface{})
```
Decode converts slice of 4 bytes to uint32. It returns number bytes consumed and
an uint32.

#### func (U32) Encode

```go
func (c U32) Encode(d interface{}) []byte
```
Encode converts uint32 to slice of 4 bytes.

#### func (U32) EncodedSize

```go
func (c U32) EncodedSize(b []byte) int
```
GetEncodedSize returns 4.

#### func (U32) Size

```go
func (c U32) Size(d interface{}) int
```
GetSize returns the size in byte after encoding v.

#### type U64

```go
type U64 struct{}
```

U64 converts uint64 to slice of 8 bytes and back.

#### func (U64) Decode

```go
func (c U64) Decode(b []byte) (int, interface{})
```
Decode converts slice of 8 bytes to uint64. It returns number bytes consumed and
an uint64.

#### func (U64) Encode

```go
func (c U64) Encode(d interface{}) []byte
```
Encode converts uint64 to slice of 8 bytes.

#### func (U64) EncodedSize

```go
func (c U64) EncodedSize(b []byte) int
```
GetEncodedSize returns 8.

#### func (U64) Size

```go
func (c U64) Size(d interface{}) int
```
GetSize returns the size in byte after encoding v.

#### type U8

```go
type U8 struct{}
```

U8 converts int8 to slice of 1 byte and back.

#### func (U8) Decode

```go
func (c U8) Decode(b []byte) (int, interface{})
```
Decode converts slice of 1 byte to int8. It returns number bytes consumed and an
int8.

#### func (U8) Encode

```go
func (c U8) Encode(d interface{}) []byte
```
Encode converts int8 to slice of 1 byte.

#### func (U8) EncodedSize

```go
func (c U8) EncodedSize(b []byte) int
```
GetEncodedSize returns 2.

#### func (U8) Size

```go
func (c U8) Size(d interface{}) int
```
GetSize returns the size in byte after encoding v.
