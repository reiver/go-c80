package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// ByteSize represents the number of bytes of a palette — i.e., how many uint8 are in a palette.
//
// Note that this is different than the number of colors in a palette.
const ByteSize = Size * c80color.ByteSize
