package c80sprite32x32

func offset(x int, y int) int {
	x = x % Width
	y = y % Height

	return y*Width + x
}
