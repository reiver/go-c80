package c80textmatrix

func (receiver Type) IsNothing() bool {
	p := receiver.bytes

	return nil == p
}
