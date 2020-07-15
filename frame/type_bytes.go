package c80frame

func (receiver *Type) Bytes() []byte {
	if nil == receiver {
		return nil
	}

	return receiver.bytes[:]
}
