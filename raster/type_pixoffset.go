package c80raster

// PixOffset returns the offset in the raster image represented by c80machine.Type that corresponds to the pixel at (x,y).
func (receiver Type) PixOffset(x, y int) int {
	return y*Width + x
}
