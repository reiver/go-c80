package c80palette

import (
	"image/color"
)

// ColorModel, in a sense, convert Type to a built-in Go color.Model.
func (receiver Type) ColorModel() color.Model {

	var p [Size]color.Color

	for index:=uint8(0); index<Size; index++ {
		p[index] = receiver.Color(index)
	}

	return color.Palette(p[:])
}
