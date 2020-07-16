package c80palettedraster

func (receiver Type) IsSomething() bool {
	return !receiver.IsNothing()
}
