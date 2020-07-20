package c80machine

import (
	"github.com/reiver/go-c80/dye"

	"github.com/reiver/go-pel"

	"image"
)

func (receiver *Type) Draw(img image.Image) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := img.(type) {
	case c80dye.Type:
		return receiver.DrawDye(casted)
	case *c80dye.Type:
		return receiver.DrawDye(casted)

	case pel.RGBA:
		return receiver.DrawPixel(casted.X, casted.Y, casted)
	case *pel.RGBA:
		return receiver.DrawPixel(casted.X, casted.Y, casted)

	default:
		return receiver.Frame().Draw(img)
	}
}
