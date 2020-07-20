package c80

import (
	"github.com/reiver/go-pel"
	"github.com/reiver/go-rgba32"

	"image"
)

func Pixel(x,y int, index uint8) image.Image {
	rgba := machine.Palette().Color(index).(rgba32.Slice)

	return pel.ColorLinked{
		X:x,
		Y:y,

		Color: rgba,
	}
}
