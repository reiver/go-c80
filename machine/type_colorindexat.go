package c80raster

// ColorIndexAt helps make c80machine.Type fit the built-in Go image.PalettedImage inteface.
//
// ColorIndexAt returns the index (into the (color) palette) of the pixel at (x,y).
func (receiver Type) ColorIndexAt(x, y int) uint8 {
	raster := receiver.Raster(0)

	return raster.ColorIndexAt(x,y)
}
