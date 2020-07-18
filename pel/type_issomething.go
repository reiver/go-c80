package c80pel

func (receiver Type) IsSomething() bool {
	return !receiver.IsNothing()
}
