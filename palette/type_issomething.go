package c80palette

func (receiver Type) IsSomething() bool {
	return !receiver.IsNothing()
}
