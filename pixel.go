package c80

import (
	"github.com/reiver/go-pel"
	"github.com/reiver/go-rgba32"

	"image"
)

func Pixel(x,y int, index uint8) image.Image {
	rgba := machine.Palette().Color(index).(rgba32.Slice)

	return pel.RGBA{
		X:x,
		Y:y,

		R: rgba[rgba32.OffsetRed],
		G: rgba[rgba32.OffsetGreen],
		B: rgba[rgba32.OffsetBlue],
		A: rgba[rgba32.OffsetAlpha],
	}
}
