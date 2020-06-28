package c80machine

import (
	"github.com/reiver/go-c80/raster"

	"image"
)

// Bounds helps make c80raster.Type fit the image.Image interface.
//
// Bounds returns the bounds for c80raster.Type.
//
// c80raster.Type is 128×192, with the top-left corner being (x,y)=(0,0),
// and the bottom-right corner being (x,y)=(127,191).
func (receiver Type) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{ X:0                   , Y:0                    ,},
		Max: image.Point{ X:(c80raster.Width-1) , Y:(c80raster.Height-1) ,},
	}
}

func (receiver Type) inBounds(x int, y int) bool {
	if c80raster.Width <= x {
		return false
	}
	if c80raster.Height <= y {
		return false
	}

	return true
}
