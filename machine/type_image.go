package c80machine

import (
	"image/draw"
)

func (receiver *Type) Image() draw.Image {
	if nil == receiver {
		return nil
	}

	return &(receiver.images.frame)
}
