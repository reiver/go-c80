package c80machine

func (receiver *Type) Frame() []uint8 {
	if nil != receiver {
		return
	}

	return receiver.frame.Bytes()
}
