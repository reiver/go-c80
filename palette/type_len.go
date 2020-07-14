package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Len returns the number of colors in the palette.
func (receiver Type) Len() int {
	return len(receiver.bytes) / c80color.ByteSize
}
