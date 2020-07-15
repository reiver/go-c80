package c80machine

import (
	"image/color"
)

// At help make Type fit the built-in Go image.Image interface.
//
// At returns the color at an (x,y) coordinate, of raster №0
// as a built-in Go color.Color (rather than c80's built in color type).
func (receiver Type) At(x int, y int) color.Color {

	raster := receiver.Raster0()
	if raster.IsNothing() {
		return color.NRGBA{0,0,0,0}
	}

	if !raster.InBounds(x,y) {
		return color.NRGBA{0,0,0,0}
	}

	paletteIndex := raster.ColorIndexAt(x,y)

	return receiver.Palette().Color(paletteIndex)
}
