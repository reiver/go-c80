package c80color

// Peek returns the color as a c80color.Array — which is just a
// [4]uint8 — in the following format:
//
//	c80color.Array{red,green,blue,alpha}
//
// I.e.,
//
//	[4]uint8{red,green,blue,alpha}
func (receiver Type) Peek() (red uint8, green uint8, blue uint8, alpha uint8) {
	p := receiver.bytes

	if nil == p {
		return 0,0,0,0
	}

	if 4 != len(p) {
		return 0,0,0,0
	}

	red   = p[offset_red]
	green = p[offset_green]
	blue  = p[offset_blue]
	alpha = p[offset_alpha]

	return red, green, blue, alpha
}
