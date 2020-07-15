package c80raster_test

import (
	"github.com/reiver/go-c80/raster"

	"math/rand"
	"time"

	"testing"
)

func TestTypeSetColorIndex(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var buffer [c80raster.Width*c80raster.Height]uint8
	var raster c80raster.Type
	var err error

	raster, err = c80raster.Wrap(buffer[:])
	if nil != err {
		t.Errorf("Received an error, but did not actually expect one.")
		t.Logf("ERROR: (%T) %q", err, err)
		return
	}

	for y := 0; y<192; y++ {
		for x := 0; x<128; x++  {

			{
				actual := raster.ColorIndexAt(x,y)

				if expected := uint8(0); expected != actual {
					t.Errorf("For test at (x,y)=(%d,%d) actual initial color index not what was expected.", x, y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}

			{
				index := uint8(randomness.Intn(16))

				raster.SetColorIndex(x,y, index)

				actual := raster.ColorIndexAt(x,y)

				if expected := index; expected != actual {
					t.Errorf("For test at (x,y)=(%d,%d) actual color index not what was expected.", x, y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}

			{
				index := uint8(randomness.Intn(16))

				raster.SetColorIndex(x,y, index)

				actual := raster.ColorIndexAt(x,y)

				if expected := index; expected != actual {
					t.Errorf("For test at (x,y)=(%d,%d) actual color index not what was expected.", x, y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}

			{
				index := uint8(2)

				raster.SetColorIndex(x,y, index)

				actual := raster.ColorIndexAt(x,y)

				if expected := index; expected != actual {
					t.Errorf("For test at (x,y)=(%d,%d) actual color index not what was expected.", x, y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}
		}
	}
}
