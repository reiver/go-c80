package c80sprite8x8

func (receiver Type) Replace(replacement ...uint8) {
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

	if 0 >= len(replacement) {
		return
	}

	copy(p[0:ByteSize], replacement)
}
