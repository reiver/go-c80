package c80palettedraster

func (receiver Type) IsNothing() bool {
	p := receiver.bytes

	return nil == p
}
