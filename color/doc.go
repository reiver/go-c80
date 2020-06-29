/*
Package c80color provides a type used for dealing with colors.

The colors are in RGBA format, where each of the 4 components is a uint8.

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

Listing

If you wanted to list out all the colors in a palette, you could do that with similar to:

	var machine c80machine.Type
	
	// ...
	
	for index:=uint8(0); index<c80palette.Size; index++ {
		
		color := machine.Palette().Color(index)
		
		fmt.Printf("Color №%d is %s", index, color)
	}
	
	// And then the output might be something such as:
	//
	// Color №0 is rgba(1,1,1,255)
	// Color №1 is rgba(222,56,43,255)
	// Color №2 is rgba(57,181,74,255)
	// Color №3 is rgba(255,199,6,255)
	// Color №4 is rgba(0,111,184,255)
	// Color №5 is rgba(118,38,113,255)
	// Color №6 is rgba(44,181,233,255)
	// Color №7 is rgba(204,204,204,255)
	// Color №8 is rgba(128,128,128,255)
	// Color №9 is rgba(255,0,0,255)
	// Color №10 is rgba(0,255,0,255)
	// Color №11 is rgba(255,255,0,255)
	// Color №12 is rgba(0,0,255,255)
	// Color №13 is rgba(255,0,255,255)
	// Color №14 is rgba(0,255,255,255)
	// Color №15 is rgba(255,255,255,255)
*/
package c80color
