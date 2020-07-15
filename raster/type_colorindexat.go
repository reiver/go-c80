package c80raster

// ColorIndexAt returns the color palette index for the pixel at (x,y).
func (receiver Type) ColorIndexAt(x, y int) uint8 {
	p := receiver.bytes

	if nil == p {
		return 0
	}
	if 0 >= len(p) {
		return 0
	}
	if ByteSize != len(p) {
		return 0
	}

	if !receiver.InBounds(x,y) {
		return 0
	}

	offset := receiver.PixOffset(x,y)

	return p[offset]
}
