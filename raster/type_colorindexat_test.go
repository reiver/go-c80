package c80raster_test

import (
	"github.com/reiver/go-c80/raster"
	"github.com/reiver/go-c80/palette"

	"testing"
)

func TestTypeColorIndexAt(t *testing.T) {

	for index:=uint8(0); index<c80palette.Size; index++ {

		testNumber := index

		var buffer [c80raster.Width*c80raster.Height]uint8

		var raster c80raster.Type = c80raster.Type(buffer[:])
		raster.Dye(index)

		for y:=0; y<c80raster.Height; y++ {
			for x:=0; x<c80raster.Width; x++ {

				actualIndex := raster.ColorIndexAt(x,y)

				if expected, actual := index, actualIndex; expected != actual {
					t.Errorf("For test #%d, actual color index for pixel (x,y)=(%d,%d) was not what was expected.", testNumber, x,y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
				}
			}
		}

		for offset,actualIndex := range raster {

			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For test #%d, actual color index for pixel at offset=%d was not what was expected.", testNumber, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
			}
		}
	}
}
