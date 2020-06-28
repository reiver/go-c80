package c80machine

// ColorIndexAt helps make c80machine.Type fit the built-in Go image.PalettedImage inteface.
//
// ColorIndexAt returns the index (into the (color) palette) of the pixel at (x,y).
func (receiver Type) ColorIndexAt(x, y int) uint8 {
	if !receiver.inBounds(x,y) {
		return 0
	}

	offset := receiver.PixOffset(x,y)

	return receiver.buffer[offset]
}
