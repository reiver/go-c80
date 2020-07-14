package c80color

// Peek returns the color split into their RGBA components.
func (receiver Type) Peek() (red uint8, green uint8, blue uint8, alpha uint8) {
	p := receiver.bytes

	if nil == p {
		return 0,0,0,0
	}

	if ByteSize != len(p) {
		return 0,0,0,0
	}

	red   = p[offset_red]
	green = p[offset_green]
	blue  = p[offset_blue]
	alpha = p[offset_alpha]

	return red, green, blue, alpha
}
