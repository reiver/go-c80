package c80raster

func (receiver Type) IsNothing() bool {
	p := receiver.bytes

	return nil == p
}
