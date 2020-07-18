package c80pel

func (receiver Type) IsNothing() bool {
	p := receiver.bytes

	return nil == p
}
