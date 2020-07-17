package c80frame

func (receiver *Type) Pix() []uint8 {
	if nil == receiver {
		return nil
	}

	return receiver.bytes[:]
}
