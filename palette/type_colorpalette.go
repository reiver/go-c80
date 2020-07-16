package c80palette

import (
	"image/color"
)

// ColorPalette returns this palette by as a Go built-in color.Palette.
func (receiver Type) ColorPalette() color.Palette {
	length := receiver.Len()

	var colorPalette color.Palette = make(color.Palette, length)

	for i:=0; i<length; i++ {
		colorPalette[i] = receiver.Color(uint8(i))
	}

	return colorPalette
}
