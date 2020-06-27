package c80raster

import (
	"github.com/reiver/go-c80/palette"
)

// Raster represents at raster image.
type Type struct {
	palette c80palette.Type
	buffer [Width*Height]byte
}
