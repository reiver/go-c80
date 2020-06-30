package c80sprite

func (receiver Type) Poke(indexes ...uint8) {
	if nil == receiver {
		return
	}

	copy(receiver, indexes)
}
