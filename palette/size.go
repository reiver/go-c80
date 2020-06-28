package c80palette

// Size represents the number of color in a palette — i.e., how many colors are in the palette.
//
// Note that this is different than the number of bytes that a palette takes up.
// For that use c80palette.Len.
//
// This is useful if you want to know how many colors there are in a palette.
// For example:
//
//	for index:=0; index<c8palette.Size; index++ {
//		
//		color := palette.Color(index)
//		
//		//@TODO
//	}
const Size = 16
