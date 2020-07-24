package c80machine

func (receiver *Type) Frame() []uint8 {
	if nil == receiver {
		return nil
	}

	return receiver.frame()
}
