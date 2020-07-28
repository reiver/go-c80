package c80

import (
	"image"
)

func Rectangle(x,y int, width, height int, index uint8) image.Image {
	return machine.Rectangle(x,y, width,height, index)
}
