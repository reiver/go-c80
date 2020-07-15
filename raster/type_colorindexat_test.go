package c80raster_test

import (
	"github.com/reiver/go-c80/raster"
	"github.com/reiver/go-c80/palette"

	"testing"
)

func TestTypeColorIndexAt(t *testing.T) {

	for index:=0; index<c80palette.Size; index++ {

		testNumber := index

		var buffer [c80raster.Width*c80raster.Height]uint8

		var raster c80raster.Type
		var err error

		raster, err = c80raster.Wrap(buffer[:])
		if nil != err {
			t.Errorf("For test #%d, received an error, but did not actually expect one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}

		raster.Dye(uint8(index))

		for y:=0; y<c80raster.Height; y++ {
			for x:=0; x<c80raster.Width; x++ {

				actualIndex := raster.ColorIndexAt(x,y)

				if expected, actual := index, int(actualIndex); expected != actual {
					t.Errorf("For test #%d, actual color index for pixel (x,y)=(%d,%d) was not what was expected.", testNumber, x,y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
				}
			}
		}
	}
}
