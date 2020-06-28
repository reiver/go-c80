package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Len represents the number of bytes of the palette — i.e., how many uint8 are in the palette.
//
// Note that this is different than the number of colors in a palette.
// For that use c80palette.Size.
//
// This is useful if you want to create a backing array to store a palette.
// For example:
//
//	var buffer [c8palette.Len]uint8
//	
//	var palette c8palette.Type = c8palette.Type(buffer[:])
const Len = Size * c80color.Len
