package c80machine

import (
	"image"
)

func (receiver *Type) Draw(img image.Image) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.Frame().Draw(img)
}
