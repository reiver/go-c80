package c80machine

// Dye changes the color of the entire raster image
// to the color represented by the (color) palette index.
func (receiver *Type) Dye(index uint8) {
	index = index & 0x0f

	for i:=0; i<(Width*Height); i++ {
		receiver.buffer[i] = index
	}
}
