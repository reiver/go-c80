package c80machine

func (receiver *Type) Terminal() Terminator {
	if nil == receiver {
		return nil
	}

	return receiver.TextMatrix()
}
