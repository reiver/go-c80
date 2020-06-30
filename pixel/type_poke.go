package c80pixel

func (receiver Type) Poke(index uint8) {
	if nil == receiver {
		return
	}
	if 0 >= len(receiver) {

	}

	receiver[0] = index
}
