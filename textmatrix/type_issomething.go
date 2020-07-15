package c80textmatrix

func (receiver Type) IsSomething() bool {
	return !receiver.IsNothing()
}
