package c80sprite8x16

func XY(offset int) (x int, y int) {
	offset = offset % Len

	x = offset % Width
	y = offset / Width

	return x, y
}
