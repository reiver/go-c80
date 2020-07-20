package c80machine

import (
	"github.com/reiver/go-pel"
	"github.com/reiver/go-rgba32"

	"image"
)

func (receiver *Type) Draw(img image.Image) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := img.(type) {
	case rgba32.Slice:
		return receiver.DrawDye(casted)
	case *rgba32.Slice:
		return receiver.DrawDye(casted)

	case pel.ColorLinked:
		return receiver.DrawPixel(casted.X, casted.Y, casted)
	case *pel.ColorLinked:
		return receiver.DrawPixel(casted.X, casted.Y, casted)

	default:
		return receiver.frame().Draw(img)
	}
}
