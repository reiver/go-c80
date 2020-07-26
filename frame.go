package c80

// Frame returns a []uint8 that represents the memory that backs
// the (main) animation frame.
//
// The bytes in this represents a series of 32-bit RGBA values,
// starting from the top-left corner, and going to the bottom-right
// corner.
func Frame() []uint8 {
	return machine.Frame()
}
