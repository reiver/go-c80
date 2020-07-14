package c80palette

func (receiver Type) Replace(c ...uint8) {
	p := receiver.bytes

	if nil == p {
		return
	}
	if 0 >= len(p) {
		return
	}

	for offset:=0; offset<len(p); offset++ {
		p[offset] = 0;
	}

	if 0 >= len(c) {
		return
	}

	copy(p[0:ByteSize], c)
}
