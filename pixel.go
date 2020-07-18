package c80

import (
	"image"
)

func Pixel(index uint8) image.Image {
	return machine.Pixel(index)
}

func DrawPixel(x int, y int, index uint8) {
	machine.DrawPixel(x,y, index)
}
