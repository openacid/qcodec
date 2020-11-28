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
