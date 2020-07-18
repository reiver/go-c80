package c80

import (
	"github.com/reiver/go-c80/pixel"

	"image"
)

func Pixel(index uint8) image.Image {
	r,g,b,a := machine.Palette().Color(index).Peek()

	return c80pixel.Type{
		R:r,
		G:g,
		B:b,
		A:a,
	}
}

func DrawPixel(x int, y int, index uint8) {
	machine.DrawPixel(x,y, index)
}
