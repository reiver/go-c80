package c80color

func RGB(red uint8, green uint8, blue uint8) Array {

	const alpha = uint8(255)

	return Array{
		red,
		green,
		blue,
		alpha,
	}
}
