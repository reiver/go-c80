package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Color returns the color in the palette at ‘index’.
//
// These color palettes only have 16 colors — from 0 to 15.
// Passing a number greater than 15 will cause it to wrap around.
func (receiver Type) Color(index uint8) c80color.Type {

	return receiver.colors[index % Size]
}
