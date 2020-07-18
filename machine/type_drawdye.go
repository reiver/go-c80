package c80machine

import (
	"image/color"
)

func (receiver *Type) DrawDye(color color.Color) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.frame.Dye(color)
}
