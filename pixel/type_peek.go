package c80pixel

func (receiver Type) Peek() (index uint8) {
	if nil == receiver {
		return 0
	}
	if 0 >= len(receiver) {
		return 0
	}

	return receiver[0]
}
