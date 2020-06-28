package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"testing"
)

func TestTypeRGBA(t *testing.T) {

	tests := []struct{
		Color c80color.Array
		ExpectedRed   uint32
		ExpectedGreen uint32
		ExpectedBlue  uint32
		ExpectedAlpha uint32
	}{
		{
			Color: c80color.Array{0,0,0, 255},
			ExpectedRed:   0x0000,
			ExpectedGreen: 0x0000,
			ExpectedBlue:  0x0000,
			ExpectedAlpha: 0xffff,
		},



		{
			Color: c80color.Array{0,0,255, 255},
			ExpectedRed:   0x0000,
			ExpectedGreen: 0x0000,
			ExpectedBlue:  0xffff,
			ExpectedAlpha: 0xffff,
		},
		{
			Color: c80color.Array{0,255,0, 255},
			ExpectedRed:   0x0000,
			ExpectedGreen: 0xffff,
			ExpectedBlue:  0x0000,
			ExpectedAlpha: 0xffff,
		},
		{
			Color: c80color.Array{0,255,255, 255},
			ExpectedRed:   0x0000,
			ExpectedGreen: 0xffff,
			ExpectedBlue:  0xffff,
			ExpectedAlpha: 0xffff,
		},
		{
			Color: c80color.Array{255,0,0, 255},
			ExpectedRed:   0xffff,
			ExpectedGreen: 0x0000,
			ExpectedBlue:  0x0000,
			ExpectedAlpha: 0xffff,
		},
		{
			Color: c80color.Array{255,0,255, 255},
			ExpectedRed:   0xffff,
			ExpectedGreen: 0x0000,
			ExpectedBlue:  0xffff,
			ExpectedAlpha: 0xffff,
		},
		{
			Color: c80color.Array{255,255,0, 255},
			ExpectedRed:   0xffff,
			ExpectedGreen: 0xffff,
			ExpectedBlue:  0x0000,
			ExpectedAlpha: 0xffff,
		},



		{
			Color: c80color.Array{255,255,255, 255},
			ExpectedRed:   0xffff,
			ExpectedGreen: 0xffff,
			ExpectedBlue:  0xffff,
			ExpectedAlpha: 0xffff,
		},
	}

	for testNumber, test := range tests {

		actualRed, actualGreen, actualBlue, actualAlpha := c80color.Type(test.Color[:]).RGBA()

		if expected, actual := test.ExpectedRed, actualRed; expected != actual {
			t.Errorf("For test #%d, the actual Red is not what was expected.", testNumber)
			t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", test.ExpectedRed, test.ExpectedGreen, test.ExpectedBlue, test.ExpectedAlpha)
			t.Logf("ACTUAL:   (r,g,b,a)=(%d,%d,%d,%d)", actualRed, actualGreen, actualBlue, actualAlpha)
			continue
		}
		if expected, actual := test.ExpectedGreen, actualGreen; expected != actual {
			t.Errorf("For test #%d, the actual Green is not what was expected.", testNumber)
			t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", test.ExpectedRed, test.ExpectedGreen, test.ExpectedBlue, test.ExpectedAlpha)
			t.Logf("ACTUAL:   (r,g,b,a)=(%d,%d,%d,%d)", actualRed, actualGreen, actualBlue, actualAlpha)
			continue
		}
		if expected, actual := test.ExpectedBlue, actualBlue; expected != actual {
			t.Errorf("For test #%d, the actual Blue is not what was expected.", testNumber)
			t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", test.ExpectedRed, test.ExpectedGreen, test.ExpectedBlue, test.ExpectedAlpha)
			t.Logf("ACTUAL:   (r,g,b,a)=(%d,%d,%d,%d)", actualRed, actualGreen, actualBlue, actualAlpha)
			continue
		}
		if expected, actual := test.ExpectedAlpha, actualAlpha; expected != actual {
			t.Errorf("For test #%d, the actual Alpha is not what was expected.", testNumber)
			t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", test.ExpectedRed, test.ExpectedGreen, test.ExpectedBlue, test.ExpectedAlpha)
			t.Logf("ACTUAL:   (r,g,b,a)=(%d,%d,%d,%d)", actualRed, actualGreen, actualBlue, actualAlpha)
			continue
		}
	}
}
