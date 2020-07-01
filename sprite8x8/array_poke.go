package c80sprite8x8

func (receiver *Array) Poke(indexes ...uint8) {
	if nil == receiver {
		return
	}

	pixel := Type(receiver[:])

	pixel.Poke(indexes...)
}
