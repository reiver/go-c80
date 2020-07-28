package c80machine

import (
	"github.com/reiver/go-rect"
	"github.com/reiver/go-rgba32"

	"image"
)

func (receiver *Type) Rectangle(x,y int, width, height int, index uint8) image.Image {
	if nil == receiver {
		return nil
	}

	rgba := receiver.Palette().Color(index).(rgba32.Slice)

	return rect.ColorLinked{
		X:x,
		Y:y,

		Width:width,
		Height:height,

		Color: rgba,
	}
}
