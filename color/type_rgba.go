package c80color

// RGBA makes Type fit the built-in Go image/color.Color interface.
func (receiver Type) RGBA() (r, g, b, a uint32) {
	const opaque = uint32(0xffff)

	return uint32(receiver.red), uint32(receiver.green), uint32(receiver.blue), opaque
}
