package c80palette

func (receiver Type) Poke(p ...uint8) {
	if nil == receiver {
		return
	}

	if nil == p {
		return
	}
	if 0 >= len(p) {
		return
	}

	copy(receiver[0:Len], p)
}
