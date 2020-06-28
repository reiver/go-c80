package c80color

// Peek return the value of the color in a [4]uint8 in the following format:
//
//	[4]uint8{red,green,blue,alpha}
func (receiver Type) Peek() (rgba [4]uint8) {
	copy(rgba[:], receiver)
	return rgba
}
