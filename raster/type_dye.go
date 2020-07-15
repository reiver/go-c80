package c80raster

// Dye changes the color of the entire raster image
// to the color represented by the (color) palette index.
func (receiver Type) Dye(index uint8) {
	p := receiver.bytes

	if nil == p {
		return
	}
	if 0 >= len(p) {
		return
	}
	if ByteSize != len(p) {
		return
	}

	limit := len(p)

	for i:=0; i<limit; i++ {
		p[i] = index
	}
}
