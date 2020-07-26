package c80

import (
	"image/color"
)

// ColorIndex returns the closest color in the color palette to the color ‘c’
// as a color index.
//
// The c80 machine uses a color palette of 256 colors.
//
// Therefore, anything that is drawn to it would only be using color's from
// its color palette.
//
// Note that the colors in the c80's color palette can be changed! If that
// happens then the index given by this previously may not be valid anymore.
// (You may have to call ColorIndex again, if you still want the closest color.)
func ColorIndex(c color.Color) uint8 {
	return machine.ColorIndex(c)
}
