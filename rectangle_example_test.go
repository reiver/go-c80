package c80_test

import (
	"github.com/reiver/go-c80"

	"fmt"
)

func ExampleRectangle() {

	// Make the entire raster image the color of color 0 in the (color) palette.
	c80.Draw(c80.Dye(0))

	// Draw something.
	//
	// What is draw is a sprite-art / pixel-art style ghost.
	{
		var left int = 78
		var top  int = 50

		var width  int = 100
		var height int = 50

		var index uint8 = 4

		c80.Draw(c80.Rectangle(left,top, width,height, index))
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
	// IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAIAAADUvDoHAAADMElEQVR4nOzVUQ0CURAEQZacOcwgBjPIAxdvc+kqBfPTmWtmHlD13B4AmwRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0q7tAef83t/tCbcxn9f2hEM8AGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkzM9sbYI0HIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmAtH8AAAD//6fMBUpntzDXAAAAAElFTkSuQmCC
}
