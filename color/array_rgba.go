package c80color

// RGBA makes Array fit the built-in Go image/color.Color interface.
func (receiver Array) RGBA() (r, g, b, a uint32) {
	return Type(receiver[:]).RGBA()
}
