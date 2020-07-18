package c80

import (
	"github.com/reiver/go-c80/dye"

	"image"
)

func Dye(index uint8) image.Image {
	r,g,b,a := machine.Palette().Color(index).Peek()

	return c80dye.Type{
		R:r,
		G:g,
		B:b,
		A:a,
	}
}
