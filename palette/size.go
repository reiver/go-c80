package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Size represents the number of color in a palette — i.e., how many colors are in the palette.
//
// Note that this is different than the number of bytes that a palette takes up.
// For that use c80palette.Len.
const Size = 16
