package c80pixel

func (receiver *Array) Poke(index uint8) {
	if nil == receiver {
		return
	}

	(*receiver)[0] = index
}
