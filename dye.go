package c80

// Dye changes the color of the entire raster image
// to the color represented by the (color) palette index.
func Dye(index uint8) {
	raster.Dye(index)
}
