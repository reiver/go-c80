package c80machine

import (
	"image/color"
)

func (receiver *Type) DrawPixel(x int, y int, color color.Color) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.Frame().Set(x,y, color)

	return nil
}
