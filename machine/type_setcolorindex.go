package c80machine

// SetColorIndex sets the color at (x,y) to the index in the palette.
func (receiver *Type) SetColorIndex(x int, y int, index uint8) {
	if nil == receiver {
		return
	}

	if !receiver.inBounds(x,y) {
		return
	}

	offset := receiver.PixOffset(x,y)

	receiver.buffer[offset] = index
}
