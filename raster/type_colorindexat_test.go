package c80raster_test

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/raster"
	"github.com/reiver/go-c80/palette"

	"math/rand"
	"time"

	"testing"
)

func TestTypeColorIndexAt(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var palette [c80palette.Size]c80color.Type
	for index:=0; index<c80palette.Size; index++ {
		palette[index] = c80color.RGB(
			uint8(randomness.Intn(256)),
			uint8(randomness.Intn(256)),
			uint8(randomness.Intn(256)),
		)
	}

	for index:=uint8(0); index<c80palette.Size; index++ {

		testNumber := index

		var raster c80raster.Type
		raster.SetPalette(c80palette.RGBs(palette[:]...))
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
	}

}
