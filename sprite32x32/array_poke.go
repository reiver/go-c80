package c80sprite32x32

func (receiver *Array) Poke(indexes ...uint8) {
	if nil == receiver {
		return
	}

	pixel := Type(receiver[:])

	pixel.Poke(indexes...)
}
