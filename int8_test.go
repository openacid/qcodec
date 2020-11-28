package qcodec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestI8(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input    int8
		want     string
		wantsize int
	}{
		{0, string([]byte{0}), 1},
		{1, string([]byte{1}), 1},
		{0x12, string([]byte{0x12}), 1},
		{^int8(0), string([]byte{0xff}), 1},
	}

	m := I8{}

	for _, c := range cases {
		rst := m.Encode(c.input)
		ta.Equal(c.want, string(rst))

		n := m.Size(c.input)
		ta.Equal(c.wantsize, n)

		n = m.EncodedSize(rst)
		ta.Equal(c.wantsize, n)

		n, u64 := m.Decode(rst)
		ta.Equal(c.input, u64)
		ta.Equal(c.wantsize, n)
	}
}
