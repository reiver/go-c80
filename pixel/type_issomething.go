package c80pixel

func (receiver Type) IsSomething() bool {
	return !receiver.IsNothing()
}
