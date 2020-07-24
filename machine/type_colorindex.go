package c80machine

import (
	"image/color"
)

func (receiver *Type) ColorIndex(c color.Color) uint8 {
	if nil == receiver {
		return 255
	}

	return receiver.Palette().Index(c)
}
