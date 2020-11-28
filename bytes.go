package qcodec

// Bytes converts a byte slice into fixed length slice.
// Result slice length is defined by Bytes.size .
type Bytes struct {
	size int
}

// Encode converts byte slice to byte slice.
func (c Bytes) Encode(d interface{}) []byte {
	return d.([]byte)
}

// Decode copies fixed length slice out of source byte slice.
// The returned bytes are NOT copied.
func (c Bytes) Decode(b []byte) (int, interface{}) {
	s := b[:c.size]
	return c.size, s
}

// GetSize returns the length: c.size.
func (c Bytes) Size(d interface{}) int {
	return c.size
}

// GetEncodedSize returns c.size
func (c Bytes) EncodedSize(b []byte) int {
	return c.size
}
