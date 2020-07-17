package c80_test

import (
	"github.com/reiver/go-c80"

	"fmt"
)

func ExampleDrawPixel() {

	// Begin
	err := c80.Begin()
	if nil != err {
		fmt.Printf("ERROR: could not begin: %s", err)
		return
	}

	// Make the entire raster image the color of color 0 in the (color) palette.
	c80.Dye(0)

	// Draw something.
	{
		left := 20
		top  := 30

		//                0
		c80.DrawPixel(left+1, top+0, 6)
		c80.DrawPixel(left+2, top+0, 6)
		c80.DrawPixel(left+3, top+0, 6)
		c80.DrawPixel(left+4, top+0, 6)
		//                5
		//                6
		//                7

		//                0
		//                1
		c80.DrawPixel(left+2, top+1, 6)
		c80.DrawPixel(left+3, top+1, 6)
		c80.DrawPixel(left+4, top+1, 6)
		c80.DrawPixel(left+5, top+1, 6)
		//                6
		//                7

		c80.DrawPixel(left+0, top+2, 6)
		//                1
		c80.DrawPixel(left+2, top+2, 6)
		c80.DrawPixel(left+3, top+2, 15)
		c80.DrawPixel(left+4, top+2, 6)
		c80.DrawPixel(left+5, top+2, 15)
		//                6
		c80.DrawPixel(left+7, top+2, 6)

		c80.DrawPixel(left+0, top+3, 6)
		//                1
		c80.DrawPixel(left+2, top+3, 6)
		c80.DrawPixel(left+3, top+3, 6)
		c80.DrawPixel(left+4, top+3, 6)
		c80.DrawPixel(left+5, top+3, 6)
		//                6
		c80.DrawPixel(left+7, top+3, 6)

		//                0
		c80.DrawPixel(left+1, top+4, 6)
		c80.DrawPixel(left+2, top+4, 6)
		c80.DrawPixel(left+3, top+4, 6)
		c80.DrawPixel(left+4, top+4, 15)
		c80.DrawPixel(left+5, top+4, 6)
		c80.DrawPixel(left+6, top+4, 6)
		//                7

		//                0
		//                1
		c80.DrawPixel(left+2, top+5, 6)
		c80.DrawPixel(left+3, top+5, 6)
		c80.DrawPixel(left+4, top+5, 6)
		c80.DrawPixel(left+5, top+5, 6)
		//                6
		//                7

		c80.DrawPixel(left+0, top+6, 6)
		//                1
		c80.DrawPixel(left+2, top+6, 6)
		c80.DrawPixel(left+3, top+6, 6)
		c80.DrawPixel(left+4, top+6, 6)
		//                5
		//                6
		//                7

		//                0
		c80.DrawPixel(left+1, top+7, 6)
		c80.DrawPixel(left+2, top+7, 6)
		c80.DrawPixel(left+3, top+7, 6)
		//                4
		//                5
		//                6
		//                7
	}

	// End
	err = c80.End()
	if nil != err {
		fmt.Printf("ERROR: could not end: %s", err)
		return
	}

	// Show the raster image.
	//
	// This will output something like:
	//
	//	IMAGE:base64-encoded-png
	//
	// Where "base64-encoded-png" is replaced with the base64 encoding of a PNG version of the raster image.
	serialized := c80.Serialize()
	fmt.Print(serialized)

	// Output:
	// IMAGE:iVBORw0KGgoAAAANSUhEUgAAAH8AAAC/BAMAAAAlex6UAAAAMFBMVEUBAQHeOCs5tUr/xwYAb7h2JnEstenMzMyAgID/AAAA/wD//wAAAP//AP8A//////9ssw/GAAAASUlEQVR4nOzRoREAMQwDwUeuM0g03T9OESIGuzw3GfkDAPaanLKQdO/PvdMFki4w+csR6g3aK0z5AQAAAAAAAAAAAFjgBQAA//8xqwgvyluaUQAAAABJRU5ErkJggg==
}
