package c80color

// Array represents an array that can be used to back a color of type c80color.Type.
//
// You might use it as:
//
//	var buffer c80color.Array
//	
//	var color c80color.Type = c80color.Type(buffer[:])
type Array [Len]uint8
