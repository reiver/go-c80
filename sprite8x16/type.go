package c80sprite8x16

// Type represents a sprite.
//
// Each sprite is a 2D grid 8×16 pixels.
//
// The way the slice values inside of c80sprite.Type correspond to the pixels is:
//
//	 0  1  2  3
//	 4  5  6  7
//	 8  9 10 11
//	12 13 14 15
//	16 17 18 19
//	20 21 22 23
//	24 25 26 27
//	28 29 30 31
type Type []uint8
