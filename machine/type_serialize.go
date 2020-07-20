package c80machine

func (receiver *Type) Serialize() string {
	return receiver.frame().String()
}
