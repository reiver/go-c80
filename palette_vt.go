package c80

import (
	"github.com/reiver/go-palette2048"
)

var palette_vt [palette2048.ByteSize]byte = [palette2048.ByteSize]byte{
	  1,  1,  1, 255, //  0 - black
	222, 56, 43, 255, //  1 - red
	 57,181, 74, 255, //  2 - green
	255,199,  6, 255, //  3 - yellow
	  0,111,184, 255, //  4 - blue
	118, 38,113, 255, //  5 - magenta
	 44,181,233, 255, //  6 - cyan
	204,204,204, 255, //  7 - white

	128,128,128, 255, //  8 - bright black
	255,  0,  0, 255, //  9 - bright red
	  0,255,  0, 255, // 10 - bright green
	255,255,  0, 255, // 11 - bright yellow
	  0,  0,255, 255, // 12 - bright blue
	255,  0,255, 255, // 13 - bright magenta
	  0,255,255, 255, // 14 - bright cyan
	255,255,255, 255, // 15 - bright white
}
