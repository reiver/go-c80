/*
Package c80color provides a type used for dealing with colors.

The colors are in RGBA format, where each of the 4 components is a uint8.
So that the size of an RGBA color is 4× the size of a uint8.

Usage

Most of the time you would probably not use c80color.Type by itself.
But instead would use a c80color.Type when using c80palette.Type
from a c80machine.Type, with code like the following:

	var machine c80machine.Type
	
	// ...
	
	// Set color in palette at ‘index’ to rgba(41,173,255, 255) — i.e., #29ADFF.
	machine.Palette().Color(index).Poke(41, 173, 255, 255)

In that code, a c80color.Type is returned from:

	machine.Palette().Color(index)
*/
package c80color
