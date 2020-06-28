package c80machine

import (
	"image/color"
)

// At help make Type fit the built-in Go image.Image interface.
//
// At returns the color at an (x,y) coordinate, of raster №0
// as a built-in Go color.Color (rather than c80's built in color type).
func (receiver Type) At(x int, y int) color.Color {

	if !receiver.inBounds(x,y) {
		return color.RGBA{}
	}

	offset := receiver.PixOffset(x,y)

	paletteIndex := receiver.Raster(0)[offset]

	return receiver.Palette().Color(paletteIndex)
}
