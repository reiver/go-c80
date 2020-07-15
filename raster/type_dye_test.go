package c80raster_test

import (
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"

	"testing"
)

func TestTypeDye(t *testing.T) {

	for index:=0; index<c80palette.Size; index++ {

		var buffer [c80raster.Width * c80raster.Height]uint8
		var raster c80raster.Type
		var err error

		raster, err = c80raster.Wrap(buffer[:])
		if nil != err {
			t.Errorf("For index=%d, received an error, but did not actually expect one.", index)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}

		raster.Dye(uint8(index))

		{
			const offset = 0

			actualIndex := buffer[offset]
			if expected, actual := index, int(actualIndex); expected != actual {
				t.Errorf("For index=%d & offset=%d, the actual index is not what was expected.", index, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 1

			actualIndex := buffer[offset]
			if expected, actual := index, int(actualIndex); expected != actual {
				t.Errorf("For index=%d & offset=%d, the actual index is not what was expected.", index, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 13

			actualIndex := buffer[offset]
			if expected, actual := index, int(actualIndex); expected != actual {
				t.Errorf("For index=%d & offset=%d, the actual index is not what was expected.", index, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 1001

			actualIndex := buffer[offset]
			if expected, actual := index, int(actualIndex); expected != actual {
				t.Errorf("For index=%d & offset=%d, the actual index is not what was expected.", index, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 10101

			actualIndex := buffer[offset]
			if expected, actual := index, int(actualIndex); expected != actual {
				t.Errorf("For index=%d & offset=%d, the actual index is not what was expected.", index, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 20202

			actualIndex := buffer[offset]
			if expected, actual := index, int(actualIndex); expected != actual {
				t.Errorf("For index=%d & offset=%d, the actual index is not what was expected.", index, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = c80raster.Width+c80raster.Height - 2

			actualIndex := buffer[offset]
			if expected, actual := index, int(actualIndex); expected != actual {
				t.Errorf("For index=%d & offset=%d, the actual index is not what was expected.", index, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = c80raster.Width+c80raster.Height - 1

			actualIndex := buffer[offset]
			if expected, actual := index, int(actualIndex); expected != actual {
				t.Errorf("For index=%d & offset=%d, the actual index is not what was expected.", index, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}


		for y:=0; y<c80raster.Height; y++ {
			for x:=0; x<c80raster.Width; x++ {
				actualIndex := raster.ColorIndexAt(x,y)

				if expected, actual := index, int(actualIndex); expected != actual {
					t.Errorf("For (x,y)=(%d,%d), the actual index is not what was expected.", x,y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}
		}
	}
}
