package c80color

// Poke sets the value of the color.
func (receiver *Type) Poke(rgba ...uint8) {
	copy(*receiver, rgba)
}
