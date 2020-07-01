package c80sprite

func Offset(x int, y int) int {
	x = x % Width
	y = y % Height

	return y*Width + x
}
