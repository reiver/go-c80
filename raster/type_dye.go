package c80raster

// Dye changes the color of the entire raster image
// to the color represented by the (color) palette index.
func (receiver Type) Dye(index uint8) {
	if nil == receiver {
		return
	}

	index = index & 0x0f

	limit := len(receiver)

	for i:=0; i<limit; i++ {
		receiver[i] = index
	}
}
