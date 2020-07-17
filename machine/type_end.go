package c80machine

func (receiver *Type) End() error {
	if nil == receiver {
		return errNilReceiver
	}

	return nil
}
