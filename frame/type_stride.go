package c80frame

func (receiver Type) Stride() int {
	return Width*Depth
}
