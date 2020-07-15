package c80palette

func (receiver Type) IsNothing() bool {
	p := receiver.bytes

	return nil == p
}
