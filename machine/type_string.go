package c80machine

func (receiver Type) String() string {
	return receiver.frame.String()
}
