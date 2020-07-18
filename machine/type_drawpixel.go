package c80machine

import (
	"image/color"
)

func (receiver *Type) DrawPixel(x int, y int, color color.Color) error {
	if nil == receiver {
		return errNilReceiver
	}

	img := receiver.Image()
	if nil == img {
		return errInternalError
	}

	img.Set(x,y, color)

	return nil
}
