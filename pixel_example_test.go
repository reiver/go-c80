package c80_test

import (
	"github.com/reiver/go-c80"

	"fmt"
)

func ExamplePixel() {

	// Make the entire raster image the color of color 0 in the (color) palette.
	c80.Draw(c80.Dye(0), 0,0)

	// Draw something.
	{
		left := 20
		top  := 30

		//                           0
		c80.Draw(c80.Pixel( 6), left+1, top+0)
		c80.Draw(c80.Pixel( 6), left+2, top+0)
		c80.Draw(c80.Pixel( 6), left+3, top+0)
		c80.Draw(c80.Pixel( 6), left+4, top+0)
		//                           5
		//                           6
		//                           7

		//                           0
		//                           1
		c80.Draw(c80.Pixel( 6), left+2, top+1)
		c80.Draw(c80.Pixel( 6), left+3, top+1)
		c80.Draw(c80.Pixel( 6), left+4, top+1)
		c80.Draw(c80.Pixel( 6), left+5, top+1)
		//                           6
		//                           7

		c80.Draw(c80.Pixel( 6), left+0, top+2)
		//                           1
		c80.Draw(c80.Pixel( 6), left+2, top+2)
		c80.Draw(c80.Pixel(15), left+3, top+2)
		c80.Draw(c80.Pixel( 6), left+4, top+2)
		c80.Draw(c80.Pixel(15), left+5, top+2)
		//                           6
		c80.Draw(c80.Pixel( 6), left+7, top+2)

		c80.Draw(c80.Pixel( 6), left+0, top+3)
		//                           1
		c80.Draw(c80.Pixel( 6), left+2, top+3)
		c80.Draw(c80.Pixel( 6), left+3, top+3)
		c80.Draw(c80.Pixel( 6), left+4, top+3)
		c80.Draw(c80.Pixel( 6), left+5, top+3)
		//                           6
		c80.Draw(c80.Pixel( 6), left+7, top+3)

		//                           0
		c80.Draw(c80.Pixel( 6), left+1, top+4)
		c80.Draw(c80.Pixel( 6), left+2, top+4)
		c80.Draw(c80.Pixel( 6), left+3, top+4)
		c80.Draw(c80.Pixel(15), left+4, top+4)
		c80.Draw(c80.Pixel( 6), left+5, top+4)
		c80.Draw(c80.Pixel( 6), left+6, top+4)
		//                           7

		//                           0
		//                           1
		c80.Draw(c80.Pixel( 6), left+2, top+5)
		c80.Draw(c80.Pixel( 6), left+3, top+5)
		c80.Draw(c80.Pixel( 6), left+4, top+5)
		c80.Draw(c80.Pixel( 6), left+5, top+5)
		//                           6
		//                           7

		c80.Draw(c80.Pixel( 6), left+0, top+6)
		//                           1
		c80.Draw(c80.Pixel( 6), left+2, top+6)
		c80.Draw(c80.Pixel( 6), left+3, top+6)
		c80.Draw(c80.Pixel( 6), left+4, top+6)
		//                           5
		//                           6
		//                           7

		//                           0
		c80.Draw(c80.Pixel( 6), left+1, top+7)
		c80.Draw(c80.Pixel( 6), left+2, top+7)
		c80.Draw(c80.Pixel( 6), left+3, top+7)
		//                           4
		//                           5
		//                           6
		//                           7
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
	// IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAIAAADUvDoHAAADbUlEQVR4nOzdwUnEUBRAUSNTg9iD2IhYrtiI2INYhYuBkP0Phsc9ZxFmM5DFu/z/mMXctm17gKrHq18AriQA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpt8Xvv3z87J+/35+X3wf+1eoJcBz6YwwwwlIA+8R/vT3tTxkwyDk7wOvn7/6EQbbFf4g5HgL3AGwCDLIagD2Y0c7ZAe5MP+OcdgUy/Ux0whUI5vJLMGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGl/AQAA//8xpiDee8Kb6AAAAABJRU5ErkJggg==
}
