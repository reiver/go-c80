package c80sprite

// Type represents a sprite.
//
// Each sprite is a 2D grid 8×8 pixels.
//
// The way the slice values inside of c80sprite.Type correspond to the pixels is:
//
//	 0  1  2  3
//	 4  5  6  7
//	 8  9 10 11
//	12 13 14 15
type Type []uint8
