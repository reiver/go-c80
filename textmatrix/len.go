package c80textmatrix

// Len represents the number of bytes of the text matrix — i.e., how many uint8 are in the text matrix.
//
// This is useful if you want to create a backing array to store a raster.
// For example:
//
//	var buffer [c8textmatrix.Len]uint8
//	
//	var textmatrix c8textmarix.Type = c8raster.Type(buffer[:])
const Len = Width * Height * Depth
