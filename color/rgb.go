package c80color

// RGB returns a color as a Type.
//
// The color is specified using separate values for ‘red’, ‘green’, and ‘blue’.
//
// Example
//
// Here is an example:
//
//	color := c80color.RGB(255,199,6)
func RGB(red uint8, green uint8, blue uint8) Type {
	alpha := uint8(0xff)

	var color [4]uint8 = [4]uint8{red,green,blue,alpha}

	return color[:]
}
