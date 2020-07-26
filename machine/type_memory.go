package c80machine

func (receiver *Type) Memory() []uint8 {
	if nil == receiver {
		return nil
	}

	return receiver.memory[:]
}
