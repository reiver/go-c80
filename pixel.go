package c80

// PixelSet sets the color of the pixel at (x,y)
// to the color represented by the index into the
// (color) palette.
func PixelSet(x int, y int, index uint8) {
	machine.Raster(0).SetColorIndex(x,y,index)
}

// PixelGet gets the color of the pixel at (x,y)
// as the color represented by the index into the
// (color) palette.
func PixelGet(x int, y int) uint8 {
	return machine.Raster(0).ColorIndexAt(x,y)
}
