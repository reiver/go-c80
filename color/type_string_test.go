package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"testing"
)

func TestTypeString(t *testing.T) {

	tests := []struct{
		Color c80color.Array
		Expected string
	}{
		{
			Color: c80color.Array{0,0,0,0},
			Expected:       "RGBA(0,0,0,0)",
		},


		{
			Color: c80color.Array{0,1,2,3},
			Expected:       "RGBA(0,1,2,3)",
		},
		{
			Color: c80color.Array{4,5,6,7},
			Expected:       "RGBA(4,5,6,7)",
		},
		{
			Color: c80color.Array{8,9,10,11},
			Expected:       "RGBA(8,9,10,11)",
		},
		{
			Color: c80color.Array{12,13,14,15},
			Expected:       "RGBA(12,13,14,15)",
		},



		{
			Color: c80color.Array{101,102,103,255},
			Expected:       "RGBA(101,102,103,255)",
		},
		{
			Color: c80color.Array{104,105,106,255},
			Expected:       "RGBA(104,105,106,255)",
		},
		{
			Color: c80color.Array{107,108,109,255},
			Expected:       "RGBA(107,108,109,255)",
		},
		{
			Color: c80color.Array{110,111,112,255},
			Expected:       "RGBA(110,111,112,255)",
		},



		{
			Color: c80color.Array{1,1,1,255},
			Expected:       "RGBA(1,1,1,255)",
		},
		{
			Color: c80color.Array{222,56,43,255},
			Expected:       "RGBA(222,56,43,255)",
		},
		{
			Color: c80color.Array{57,181,74,255},
			Expected:       "RGBA(57,181,74,255)",
		},
		{
			Color: c80color.Array{255,199,6,255},
			Expected:       "RGBA(255,199,6,255)",
		},
		{
			Color: c80color.Array{0,111,184,255},
			Expected:       "RGBA(0,111,184,255)",
		},
		{
			Color: c80color.Array{118,38,113,255},
			Expected:       "RGBA(118,38,113,255)",
		},
		{
			Color: c80color.Array{44,181,233,255},
			Expected:       "RGBA(44,181,233,255)",
		},
		{
			Color: c80color.Array{204,204,204,255},
			Expected:       "RGBA(204,204,204,255)",
		},



		{
			Color: c80color.Array{255,255,255,255},
			Expected:       "RGBA(255,255,255,255)",
		},
	}

	for testNumber, test := range tests {

		actual := c80color.Type(test.Color[:]).String()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
