package c80machine

import (
	"github.com/reiver/go-c80/palette"
)

// Type represents a virtual machine that has a raster image, and a palette.
type Type struct {
	palette c80palette.Type
	buffer [Width*Height]byte
}
