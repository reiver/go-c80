package c80machine

func (receiver *Type) FrameBytes() []byte {
	if nil == receiver {
		return nil
	}

	return receiver.frame.Bytes()
}
