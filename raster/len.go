package c80raster

// Len represents the number of bytes of the raster — i.e., how many uint8 are in the raster.
//
// This is useful if you want to create a backing array to store a raster.
// For example:
//
//	var buffer [c8raster.Len]uint8
//	
//	var raster c8raster.Type = c8raster.Type(buffer[:])
const Len = Width * Height
