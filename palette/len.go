package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Len represents the number of bytes of the palette — i.e., how many uint8 are in the palette.
//
// Note that this is different than the number of colors in a palette.
// For that use c80palette.Size.
const Len = Size * c80color.Len
