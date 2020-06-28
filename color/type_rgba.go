package c80color

// RGBA makes Type fit the built-in Go image/color.Color interface.
func (receiver Type) RGBA() (r, g, b, a uint32) {
	const opaque = uint32(0xffff)

	var red16   uint32 = uint32(receiver.red)   * (0xffff/0xff)
	var green16 uint32 = uint32(receiver.green) * (0xffff/0xff)
	var blue16  uint32 = uint32(receiver.blue)  * (0xffff/0xff)

	return red16, green16, blue16, opaque
}
