package c80machine

import (
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"
)

const NumberOfRasters = 3

const Len = c80palette.Len + NumberOfRasters*c80raster.Len

// Type represents a virtual machine that has a raster image, and a palette.
type Type struct {
	memory [Len]uint8
}

func (receiver Type) Palette() c80palette.Type {
	beginning := 0
	ending    := beginning + c80palette.Len

	return receiver.memory[beginning:ending]
}

func (receiver Type) Raster(number int) c80raster.Type {
	number = number % NumberOfRasters

	beginning := c80palette.Len + number*c80raster.Len
	ending    := beginning + c80raster.Len

	return receiver.memory[beginning:ending]
}
