package c80sprite8x8

func (receiver Type) Poke(indexes ...uint8) {
	if nil == receiver {
		return
	}

	copy(receiver, indexes)
}