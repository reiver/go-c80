package c80sprite8x8

func offset(x int, y int) int {
	x = x % Width
	y = y % Height

	return y*Width + x
}
