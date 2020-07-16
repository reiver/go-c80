package c80pixel

func (receiver Type) IsNothing() bool {
	p := receiver.bytes

	return nil == p
}
