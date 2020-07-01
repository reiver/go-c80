package c80sprite

func XY(offset int) (x int, y int) {
	offset = offset % Len

	x = offset % Width
	y = offset / Height

	return x, y
}
