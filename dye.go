package c80

import (
	"github.com/reiver/go-c80/dye"

	"github.com/reiver/go-rgba32"

	"image"
)

func Dye(index uint8) image.Image {
	rgba := machine.Palette().Color(index).(rgba32.Slice)

	return c80dye.Type{
		R: rgba[rgba32.OffsetRed],
		G: rgba[rgba32.OffsetGreen],
		B: rgba[rgba32.OffsetBlue],
		A: rgba[rgba32.OffsetAlpha],
	}
}
