package c80color

func (receiver Type) IsSomething() bool {
	return !receiver.IsNothing()
}
