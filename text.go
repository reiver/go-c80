package c80

import (
	"image"
)

func Text(msg string) image.Image {
	return machine.Text(msg)
}

func DrawText(x int, y int, msg string) {
	machine.DrawText(x,y, msg)
}
