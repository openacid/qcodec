package qcodec

import "testing"

var Output int

func BenchmarkI32(b *testing.B) {
	c := I32{}

	s := uint8(0)
	for i := 0; i < b.N; i++ {
		bs := c.Encode(int32(3))

		s += bs[0]
	}

	Output = int(s)

}
