package c80_test

import (
	"github.com/reiver/go-c80"

	"fmt"
)

func ExampleDraw() {

	// Set the color palette to "vt".
	err := c80.Colorize("vt")
	if nil != err {
		fmt.Printf("ERROR: could not set the palette: %s", err)
		return
	}

	const bgIndex = 4 // 4 is blue  in the "vt" color palette.
	const fgIndex = 7 // 7 is white in the "vt" color palette.


	// Clear the whole frame with the backgroud color.
	//
	// That means that every pixel of the frame will have the same color.
	err = c80.Draw(c80.Dye(bgIndex))
	if nil != err {
		fmt.Printf("ERROR: could not clear the image with color at index=%d: %s", bgIndex, err)
		return
	}


	const left = 5
	const top  = 10

	// Here we draw something that looks like a key.
	//
	//	▄▄▄▄▄█▀█
	//	▀ ▀  ▀▀▀
	{
		//                      0
		//                      1
		//                      2
		//                      3
		//                      4
		c80.Draw(c80.Pixel(left+5, top+0, fgIndex))
		c80.Draw(c80.Pixel(left+6, top+0, fgIndex))
		c80.Draw(c80.Pixel(left+7, top+0, fgIndex))

		c80.Draw(c80.Pixel(left+0, top+1, fgIndex))
		c80.Draw(c80.Pixel(left+1, top+1, fgIndex))
		c80.Draw(c80.Pixel(left+2, top+1, fgIndex))
		c80.Draw(c80.Pixel(left+3, top+1, fgIndex))
		c80.Draw(c80.Pixel(left+4, top+1, fgIndex))
		c80.Draw(c80.Pixel(left+5, top+1, fgIndex))
		//                      6
		c80.Draw(c80.Pixel(left+7, top+1, fgIndex))

		c80.Draw(c80.Pixel(left+0, top+2, fgIndex))
		//                      1
		c80.Draw(c80.Pixel(left+2, top+2, fgIndex))
		//                      3
		//                      4
		c80.Draw(c80.Pixel(left+5, top+2, fgIndex))
		c80.Draw(c80.Pixel(left+6, top+2, fgIndex))
		c80.Draw(c80.Pixel(left+7, top+2, fgIndex))
	}

	// Here we output a serialized version of the frame.
	//
	// This will output something like:
	//
	//	IMAGE:base64-encoded-png
	//
	// Where "base64-encoded-png" is replaced with the base64 encoding of a PNG version of the raster image.
	fmt.Print(c80.Serialize())

	// Output:
	//
	// IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAIAAADUvDoHAAADTklEQVR4nOzasQ2CUBRAUTSM4Cis4wwOxRIO4SgMYUFCocYSEu451Q80r3iXQMI4PJ4DVF2PHgCOJADSBECaAEgTAGkCIE0ApAmANAGQNv6//brf1sM0L7vMA7v6DGDb+M00L98X4Rx+vAKtD/tpXtaD7efELn6Go8xHMGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIO0dAAD//1sjEleiMSMyAAAAAElFTkSuQmCC
}
