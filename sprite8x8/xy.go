package c80sprite8x8

func XY(offset int) (x int, y int) {
	offset = offset % Len

	x = offset % Width
	y = offset / Height

	return x, y
}
