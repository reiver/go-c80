package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"testing"
)

func TestTypeString(t *testing.T) {

	tests := []struct{
		Color [4]uint8
		Expected string
	}{
		{
			Color: [4]uint8{0,0,0,0},
			Expected: "rgba(0,0,0,0)",
		},


		{
			Color: [4]uint8{0,1,2,3},
			Expected: "rgba(0,1,2,3)",
		},
		{
			Color: [4]uint8{4,5,6,7},
			Expected: "rgba(4,5,6,7)",
		},
		{
			Color: [4]uint8{8,9,10,11},
			Expected: "rgba(8,9,10,11)",
		},
		{
			Color: [4]uint8{12,13,14,15},
			Expected: "rgba(12,13,14,15)",
		},



		{
			Color: [4]uint8{101,102,103,255},
			Expected: "rgba(101,102,103,255)",
		},
		{
			Color: [4]uint8{104,105,106,255},
			Expected: "rgba(104,105,106,255)",
		},
		{
			Color: [4]uint8{107,108,109,255},
			Expected: "rgba(107,108,109,255)",
		},
		{
			Color: [4]uint8{110,111,112,255},
			Expected: "rgba(110,111,112,255)",
		},



		{
			Color: [4]uint8{1,1,1,255},
			Expected: "rgba(1,1,1,255)",
		},
		{
			Color: [4]uint8{222,56,43,255},
			Expected: "rgba(222,56,43,255)",
		},
		{
			Color: [4]uint8{57,181,74,255},
			Expected: "rgba(57,181,74,255)",
		},
		{
			Color: [4]uint8{255,199,6,255},
			Expected: "rgba(255,199,6,255)",
		},
		{
			Color: [4]uint8{0,111,184,255},
			Expected: "rgba(0,111,184,255)",
		},
		{
			Color: [4]uint8{118,38,113,255},
			Expected: "rgba(118,38,113,255)",
		},
		{
			Color: [4]uint8{44,181,233,255},
			Expected: "rgba(44,181,233,255)",
		},
		{
			Color: [4]uint8{204,204,204,255},
			Expected: "rgba(204,204,204,255)",
		},



		{
			Color: [4]uint8{255,255,255,255},
			Expected: "rgba(255,255,255,255)",
		},
	}

	for testNumber, test := range tests {

		color, err := c80color.Wrap(test.Color[:])
		if nil != err {
			t.Errorf("For test #%d, received an error, but did not actually expect one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}

		actual := color.String()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
