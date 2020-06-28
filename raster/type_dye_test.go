package c80raster_test

import (
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"

	"testing"
)

func TestTypeDye(t *testing.T) {

	for index:=uint8(0); index<c80palette.Size; index++ {

		var buffer [c80raster.Width * c80raster.Height]uint8

		var raster c80raster.Type = c80raster.Type(buffer[:])

		raster.Dye(index)

		{
			const offset = 0

			actualIndex := buffer[offset]
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 1

			actualIndex := buffer[offset]
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 13

			actualIndex := buffer[offset]
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 1001

			actualIndex := buffer[offset]
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 10101

			actualIndex := buffer[offset]
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = 20202

			actualIndex := buffer[offset]
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = c80raster.Width+c80raster.Height - 2

			actualIndex := buffer[offset]
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			const offset = c80raster.Width+c80raster.Height - 1

			actualIndex := buffer[offset]
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}


		for offset,actualIndex := range raster {
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}
	}
}
