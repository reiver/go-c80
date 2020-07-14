package c80color

func (receiver Type) IsNothing() bool {
	p := receiver.bytes

	return nil == p
}
