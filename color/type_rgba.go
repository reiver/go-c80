package c80color

// RGBA makes Type fit the built-in Go image/color.Color interface.
func (receiver Type) RGBA() (r, g, b, a uint32) {
	if nil == receiver {
		return 0,0,0,0
	}
	if 0 >= len(receiver) {
		return 0,0,0,0
	}

	var red16    uint32
	var green16  uint32
	var blue16   uint32
	var alpha16  uint32

	if 0 < len(receiver) {
		red16    = uint32(receiver[0]) * (0xffff/0xff)
	}
	if 1 < len(receiver) {
		green16  = uint32(receiver[1]) * (0xffff/0xff)
	}
	if 2 < len(receiver) {
		blue16   = uint32(receiver[2]) * (0xffff/0xff)
	}
	if 3 < len(receiver) {
		alpha16  = uint32(receiver[3]) * (0xffff/0xff)
	}

	return red16, green16, blue16, alpha16
}
