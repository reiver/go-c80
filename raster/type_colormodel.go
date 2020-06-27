package c80raster

import (
	"image/color"
)

// Bounds helps make c80raster.Type fit the image.Image interface.
func (receiver Type) ColorModel() color.Model {
	return receiver.palette.Model()
}
