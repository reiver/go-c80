package c80raster

import (
	"github.com/reiver/go-c80/pixel"
)

// ByteSize represents the number of bytes of the raster — i.e., how many uint8 are in the raster.
const ByteSize = Width * Height * c80pixel.ByteSize