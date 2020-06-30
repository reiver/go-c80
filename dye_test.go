package c80_test

import (
	"github.com/reiver/go-c80"
	"github.com/reiver/go-c80/machine"
	"github.com/reiver/go-c80/palette"

	"testing"
)

func TestDye(t *testing.T) {

	for index:=uint8(0); index<c80palette.Size; index++ {

		testNumber := index

		c80.Dye(index)

		for y:=0; y<c80raster.Height; y++ {
			for x:=0; x<c80raster.Width; x++ {
				actual := c80.PixelGet(x,y)

				if expected := index; expected != actual {
					t.Errorf("For test #%d, actual color index at (x,y)=(%d,%d) is not what was expected.", testNumber, x,y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}
		}
	}
}
