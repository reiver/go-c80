package c80machine

import (
	"image/color"
)

// Bounds helps make c80machine.Type fit the image.Image interface.
func (receiver Type) ColorModel() color.Model {
	return receiver.palette.ColorModel()
}
