package c80_test

/*
import (
	"github.com/reiver/go-c80"
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"

	"bytes"
	"image"
	"image/png"

	"testing"
)

func TestDye(t *testing.T) {

	c80.Colorize("tia")

	Loop: for index:=0; index<c80palette.Len; index++ {

		testNumber := index

		c80.Dye(uint8(index))

		var pngBuffer bytes.Buffer
		if err := c80.PNGTo(&pngBuffer); nil != err {
			t.Errorf("For test #%d, received an error, but did not actually expect one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}

		var img image.Image
		{
			var err error

			img, err = png.Decode(&pngBuffer)
			if nil != err {
				t.Errorf("For test #%d, received an error, but did not actually expect one.", testNumber)
				t.Logf("ERROR: (%T) %q", err, err)
				continue
			}
		}

		var eR, eG, eB, eA uint32 = 4,5,6,7
		{
			c80.
		}

		for y:=0; y<c80raster.Height; y++ {
			for x:=0; x<c80raster.Width; x++ {
				color := img.At(x,y)

				aR, aG, aB, aA := color.RGBA()

				if eR != aR || eG != aG || eB != aB || eA != aA {
					t.Errorf("For test #%d, actual color at (x,y)=(%d,%d) is not what was expected.", testNumber, x,y)
					t.Logf("COLOR INDEX: %d", index)
					t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, eA)
					t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
					continue Loop
				}
			}
		}
	}
}
*/
