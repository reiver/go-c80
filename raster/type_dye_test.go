package c80raster

import (
	"github.com/reiver/go-c80/palette"

	"testing"
)

func TestTypeDye(t *testing.T) {

	for index:=uint8(0); index<c80palette.Size; index++ {

		var raster Type

		raster.Dye(index)

		for offset,actualIndex := range raster.buffer {
			if expected, actual := index, actualIndex; expected != actual {
				t.Errorf("For offset=%d, the actual index is not what was expected.", offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}
	}
}
