package c80

// PixelSet sets the color of the pixel at (x,y)
// to the color represented by the index into the
// (color) palette.
func PixelSet(x int, y int, index uint8) {

	raster.SetColorIndex(x,y,index)
}

func PixelGet(x int, y int) uint8 {
	return raster.ColorIndexAt(x,y)
}
