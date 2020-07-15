package c80pixel

func (receiver Type) Poke(index uint8) {
	p := receiver.bytes

	if nil == p {
		return
	}
	if 0 >= len(p) {
		return
	}
	if ByteSize != len(p) {
		return
	}

	p[0] = index
}
