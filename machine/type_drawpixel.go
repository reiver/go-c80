package c80machine

import (
	"image/color"
)

func (receiver *Type) DrawPixel(x int, y int, index uint8) {
	if nil == receiver {
		return
	}

	var color color.Color
	{
		palette := receiver.Palette()
		if palette.IsNothing() {
			return
		}

		color = palette.Color(index)
	}

	image := receiver.Image()
	if nil == receiver {
		return
	}

	image.Set(x,y, color)
}
