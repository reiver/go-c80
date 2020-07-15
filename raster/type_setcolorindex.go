package c80raster

// SetColorIndex sets the color at (x,y) to the index in the palette.
func (receiver Type) SetColorIndex(x int, y int, index uint8) {
	p := receiver.bytes

	if nil == p {
		return
	}
	if 0 >= len(p) {
		return
	}
	if ByteSize != len(p) {
		return
	}

	if !receiver.InBounds(x,y) {
		return
	}

	offset := receiver.PixOffset(x,y)

	p[offset] = index
}
