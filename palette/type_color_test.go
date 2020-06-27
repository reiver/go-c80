package c80palette_test

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/palette"

	"testing"
)

func TestTypeColor(t *testing.T) {

	tests := []struct{
		Palette [c80palette.Size]c80color.Type
	}{
		{
			// Nothing here.
		},
		{
			Palette: [c80palette.Size]c80color.Type{
				c80color.RGB( 0, 1, 2),
				c80color.RGB( 3, 4, 5),
				c80color.RGB( 6, 7, 8),
				c80color.RGB( 9,10,11),

				c80color.RGB(12,13,14),
				c80color.RGB(15,16,17),
				c80color.RGB(18,19,20),
				c80color.RGB(21,22,23),

				c80color.RGB(24,25,26),
				c80color.RGB(27,28,29),
				c80color.RGB(30,31,32),
				c80color.RGB(33,34,35),

				c80color.RGB(36,37,38),
				c80color.RGB(39,40,41),
				c80color.RGB(42,43,44),
				c80color.RGB(45,46,47),
			},
		},
		{
			Palette: [c80palette.Size]c80color.Type{
				c80color.RGB(  1,  1,  1),
				c80color.RGB(222, 56, 43),
				c80color.RGB( 57,181, 74),
				c80color.RGB(255,199,  6),

				c80color.RGB(  0,111,184),
				c80color.RGB(118, 38,113),
				c80color.RGB( 44,181,233),
				c80color.RGB(204,204,204),

				c80color.RGB(128,128,128),
				c80color.RGB(255,  0,  0),
				c80color.RGB(  0,255,  0),
				c80color.RGB(255,255,  0),

				c80color.RGB(  0,  0,255),
				c80color.RGB(255,  0,255),
				c80color.RGB(  0,255,255),
				c80color.RGB(255,255,255),
			},
		},
	}

	for testNumber, test := range tests {

		var palette c80palette.Type = c80palette.RGBs(test.Palette[:]...)

		for index:=uint8(0); index<c80palette.Size; index++ {
			expected := test.Palette[index]
			actual := palette.Color(index)

			if expected != actual {
				t.Errorf("For test #%d, the actual index is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}
	}
}
