package c80sprite32x32

func (receiver Type) Poke(indexes ...uint8) {
	if nil == receiver {
		return
	}

	copy(receiver, indexes)
}
