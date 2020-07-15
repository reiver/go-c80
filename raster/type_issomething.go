package c80raster

func (receiver Type) IsSomething() bool {
	return !receiver.IsNothing()
}
