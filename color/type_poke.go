package c80color

// Poke sets the value of the color.
func (receiver Type) Poke(red uint8, green uint8, blue uint8, alpha uint8) {
	p := receiver.bytes

	if nil == p {
		return
	}

	if ByteSize != len(p) {
		return
	}

	p[offset_red]   = red
	p[offset_green] = green
	p[offset_blue]  = blue
	p[offset_alpha] = alpha
}
