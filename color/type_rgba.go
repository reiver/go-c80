package c80color

// RGBA makes Type fit the built-in Go image/color.Color interface.
func (receiver Type) RGBA() (r, g, b, a uint32) {
	p := receiver.bytes

	if nil == p {
		return 0,0,0,0
	}
	if 0 >= len(p) {
		return 0,0,0,0
	}

	if ByteSize != len(p) {
		return 0,0,0,0
	}

	var red16    uint32
	var green16  uint32
	var blue16   uint32
	var alpha16  uint32

	red16   = uint32(p[offset_red])   * (0xffff/0xff)
	green16 = uint32(p[offset_green]) * (0xffff/0xff)
	blue16  = uint32(p[offset_blue])  * (0xffff/0xff)
	alpha16 = uint32(p[offset_alpha]) * (0xffff/0xff)

	return red16, green16, blue16, alpha16
}
