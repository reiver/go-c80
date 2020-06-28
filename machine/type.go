package c80machine

import (
	"github.com/reiver/go-c80/palette"
)

const Len = c80palette.Len

// Type represents a virtual machine that has a raster image, and a palette.
type Type struct {
	memory [Len]uint8
}

func (receiver Type) Palette() c80palette.Type {
	beginning := 0
	ending    := beginning + c80palette.Len

	return memory[beginning:ending]
}
