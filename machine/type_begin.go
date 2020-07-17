package c80machine

func (receiver *Type) Begin() error {
	if nil == receiver {
		return errNilReceiver
	}

	return nil
}
