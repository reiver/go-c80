package c80color

// Peek returns the color as a c80color.Array — which is just a
// [4]uint8 — in the following format:
//
//	c80color.Array{red,green,blue,alpha}
//
// I.e.,
//
//	[4]uint8{red,green,blue,alpha}
func (receiver Type) Peek() (rgba Array) {
	copy(rgba[:], receiver)
	return rgba
}
