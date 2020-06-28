package c80raster

// ColorIndexAt returns the color palette index for the pixel at (x,y).
func (receiver Type) ColorIndexAt(x, y int) uint8 {
	if !receiver.InBounds(x,y) {
		return 0
	}

	offset := receiver.PixOffset(x,y)

	return receiver[offset]
}
