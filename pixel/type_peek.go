package c80pixel

func (receiver Type) Peek() (index uint8) {
	p := receiver.bytes

	if nil == p {
		return 0
	}
	if 0 >= len(p) {
		return 0
	}
	if ByteSize != len(p) {
		return 0
	}

	return p[0]
}
