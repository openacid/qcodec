package qcodec

import (
	"reflect"
	"testing"
)

func TestBytes(t *testing.T) {

	var x Codec = Bytes{}
	_ = x

	cases := []struct {
		input []byte
		want  int
	}{
		{[]byte(""), 0},
		{[]byte("a"), 1},
		{[]byte("abc"), 3},
	}

	for i, c := range cases {
		m := Bytes{size: c.want}
		l := m.Size(c.input)
		if l != c.want {
			t.Fatalf("%d-th: Size: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, l)
		}

		rst := m.Encode(c.input)
		if len(rst) != c.want {
			t.Fatalf("%d-th: encoded len: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, len(rst))
		}

		l = m.EncodedSize(rst)
		if l != c.want {
			t.Fatalf("%d-th: encoded size: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, l)
		}

		n, s := m.Decode(rst)
		if c.want != n {
			t.Fatalf("%d-th: decoded size: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, n)
		}
		if !reflect.DeepEqual(c.input, s) {
			t.Fatalf("%d-th: decode: input: %v; want: %v; actual: %v",
				i+1, c.input, c.input, s)
		}

		if len(rst) > 0 {
			rst[0] = 'x'
			if s.([]byte)[0] != 'x' {
				t.Fatalf("should be not be copied.")
			}
		}

	}
}
