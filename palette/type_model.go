package c80palette

import (
	"image/color"
)

// Model, in a sense, convert Type to a built-in Go color.Model.
func (receiver Type) Model() color.Model {

	var p [Size]color.Color

	for index:=uint8(0); index<Size; index++ {
		p[index] = receiver.Color(index)
	}

	return color.Palette(p[:])
}
