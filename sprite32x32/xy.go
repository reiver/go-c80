package c80sprite32x32

func XY(offset int) (x int, y int) {
	offset = offset % Len

	x = offset % Width
	y = offset / Width

	return x, y
}
