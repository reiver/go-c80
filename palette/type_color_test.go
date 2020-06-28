package c80palette_test

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/palette"

	"testing"
)

func TestTypeColor(t *testing.T) {

	tests := []struct{
		Palette [c80palette.Len]uint8
	}{
		{
			// Nothing here.
		},
		{
			Palette: [c80palette.Len]uint8{
				 0, 1, 2, 3,
				 4, 5, 6, 7,
				 8, 9,10,11,
				12,13,14,15,

				16,17,18,19,
				20,21,22,23,
				24,25,26,27,
				28,29,30,31,

				32,33,34,35,
				36,37,38,39,
				40,41,42,43,
				44,45,46,47,

				48,49,50,51,
				52,53,54,55,
				56,57,58,59,
				60,61,62,63,
			},
		},
		{
			Palette: [c80palette.Len]uint8{
				  1,  1,  1,255,
				222, 56, 43,255,
				 57,181, 74,255,
				255,199,  6,255,

				  0,111,184,255,
				118, 38,113,255,
				 44,181,233,255,
				204,204,204,255,

				128,128,128,255,
				255,  0,  0,255,
				  0,255,  0,255,
				255,255,  0,255,

				  0,  0,255,255,
				255,  0,255,255,
				  0,255,255,255,
				255,255,255,255,
			},
		},
	}

	for testNumber, test := range tests {

		var palette c80palette.Type = c80palette.Type(test.Palette[:])

		for index:=uint8(0); index<c80palette.Size; index++ {
			var expected [c80color.Len]uint8
			for i:=uint8(0); i<c80color.Len; i++ {
				expected[i] = test.Palette[index*c80color.Len + i]
			}

			actual := palette.Color(index)

			if c80color.Array(expected).String() != actual.String() {
				t.Errorf("For test #%d, the actual color is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}
	}
}
