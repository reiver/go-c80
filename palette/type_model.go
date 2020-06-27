package c80palette

import (
	"image/color"
)

// Model, in a sense, convert Type to a built-in Go color.Model.
func (receiver Type) Model() color.Model {

	var p [Size]color.Color

	for i, color := range receiver.colors {
		p[i] = color
	}

	return color.Palette(p[:])
}
