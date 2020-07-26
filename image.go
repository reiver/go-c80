package c80

import(
	"image/draw"
)

// Image return a draw.Image for the (main) animation frame.
func Image() draw.Image {
	return machine.Image()
}
