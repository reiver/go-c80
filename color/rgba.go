package c80color

func RGBA(red uint8, green uint8, blue uint8, alpha uint8) Array {

	return Array{
		red,
		green,
		blue,
		alpha,
	}
}
