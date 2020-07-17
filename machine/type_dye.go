package c80machine

func (receiver *Type) Dye(index uint8) {
	if nil == receiver {
		return
	}

	palette := receiver.Palette()
	if palette.IsNothing() {
		return
	}

	color := palette.Color(index)
	if color.IsNothing() {
		return
	}

	receiver.frame.Dye(color)
}
