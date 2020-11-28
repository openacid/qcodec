package qcodec

// Dummy converts anything to nothing.
type Dummy struct {
	size int
}

// Encode converts something to empty byte slice.
func (c Dummy) Encode(d interface{}) []byte {
	return []byte{}
}

// Decode always returns nil.
func (c Dummy) Decode(b []byte) (int, interface{}) {
	return 0, nil
}

// GetSize returns 0.
func (c Dummy) Size(d interface{}) int {
	return 0
}

// GetEncodedSize returns 0.
func (c Dummy) EncodedSize(b []byte) int {
	return 0
}
