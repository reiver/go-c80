package c80machine

import (
	"image"
)

// Bounds helps make c80machine.Type fit the image.Image interface.
//
// Bounds returns the bounds for c80machine.Type.
//
// c80machine.Type is 128×192, with the top-left corner being (x,y)=(0,0),
// and the bottom-right corner being (x,y)=(127,191).
func (receiver Type) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{ X:0         , Y:0          ,},
		Max: image.Point{ X:(Width-1) , Y:(Height-1) ,},
	}
}

func (receiver Type) inBounds(x int, y int) bool {
	if Width <= x {
		return false
	}
	if Height <= y {
		return false
	}

	return true
}
