package c80machine

import (
	"image"
	"image/color"
	"image/draw"
)

func (receiver *Type) Image() draw.Image {
	if nil == receiver {
		return nil
	}

	return receiver
}

func (receiver *Type) At(x, y int) color.Color {
	if nil == receiver {
		return nil
	}

	return receiver.frame().At(x,y)
}

func (receiver *Type) Bounds() image.Rectangle {
	if nil == receiver {
		return image.Rectangle{}
	}

	return receiver.frame().Bounds()
}

func (receiver *Type) ColorModel() color.Model {
	if nil == receiver {
		return nil
	}

	return receiver.frame().ColorModel()
}

func (receiver *Type) Set(x, y int, c color.Color) {
	if nil == receiver {
		return
	}

	receiver.frame().Set(x,y, c)
}
