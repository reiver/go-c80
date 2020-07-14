package c80palette

func (receiver Type) Replace(p ...uint8) {
	if nil == receiver {
		return
	}

	if nil == p {
		return
	}
	if 0 >= len(p) {
		return
	}

	for offset:=0; offset<len(receiver); offset++ {
		receiver[offset] = 0;
	}

	copy(receiver[0:Len], p)

}
