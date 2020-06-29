package c80machine

import (
	"image/color"
)

// ColorModel helps make c80machine.Type fit the image.Image interface.
func (receiver Type) ColorModel() color.Model {
	return receiver.Palette().ColorModel()
}
