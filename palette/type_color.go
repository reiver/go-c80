package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Color returns the color in the palette at ‘index’.
//
// These color palettes only have 64 colors — from 0 to 63.
// Passing a number greater than 63 will cause it to wrap around.
//
// (So, for example, asking for color ‘66’ would give you color ‘2’.)
func (receiver Type) Color(index uint8) c80color.Type {
	index = index % Size

	beginning := index * c80color.Len
	ending := beginning + c80color.Len

	return c80color.Type(receiver[beginning:ending])
}
