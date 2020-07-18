package c80dye

import (
	"image"
	"image/color"
)

type Type struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (receiver Type) At(x, y int) color.Color {
	return receiver
}

func (receiver Type) Bounds() image.Rectangle {
	const max = int((^uint(0)) >> 1)
	const min = -max - 1

	return image.Rectangle{
		Min: image.Point{min,min},
		Max: image.Point{max,max},
	}
}

func (receiver Type) ColorModel() color.Model {
	return color.RGBAModel
}

func (receiver Type) RGBA() (r, g, b, a uint32) {
	r = uint32(receiver.R) * (0xffff/0xff)
	g = uint32(receiver.G) * (0xffff/0xff)
	b = uint32(receiver.B) * (0xffff/0xff)
	a = uint32(receiver.A) * (0xffff/0xff)

	return r,g,b,a
}
