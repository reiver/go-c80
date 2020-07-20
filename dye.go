package c80

import (
	"github.com/reiver/go-rgba32"

	"image"
)

func Dye(index uint8) image.Image {
	rgba := machine.Palette().Color(index).(rgba32.Slice)

	return rgba
}
