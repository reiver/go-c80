package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"testing"
)

func TestArrayString(t *testing.T) {

	tests := []struct{
		Color c80color.Array
		Expected string
	}{
		{
			Color: c80color.Array{0,0,0,0},
			Expected:       "rgba(0,0,0,0)",
		},


		{
			Color: c80color.Array{0,1,2,3},
			Expected:       "rgba(0,1,2,3)",
		},
		{
			Color: c80color.Array{4,5,6,7},
			Expected:       "rgba(4,5,6,7)",
		},
		{
			Color: c80color.Array{8,9,10,11},
			Expected:       "rgba(8,9,10,11)",
		},
		{
			Color: c80color.Array{12,13,14,15},
			Expected:       "rgba(12,13,14,15)",
		},



		{
			Color: c80color.Array{101,102,103,255},
			Expected:       "rgba(101,102,103,255)",
		},
		{
			Color: c80color.Array{104,105,106,255},
			Expected:       "rgba(104,105,106,255)",
		},
		{
			Color: c80color.Array{107,108,109,255},
			Expected:       "rgba(107,108,109,255)",
		},
		{
			Color: c80color.Array{110,111,112,255},
			Expected:       "rgba(110,111,112,255)",
		},



		{
			Color: c80color.Array{1,1,1,255},
			Expected:       "rgba(1,1,1,255)",
		},
		{
			Color: c80color.Array{222,56,43,255},
			Expected:       "rgba(222,56,43,255)",
		},
		{
			Color: c80color.Array{57,181,74,255},
			Expected:       "rgba(57,181,74,255)",
		},
		{
			Color: c80color.Array{255,199,6,255},
			Expected:       "rgba(255,199,6,255)",
		},
		{
			Color: c80color.Array{0,111,184,255},
			Expected:       "rgba(0,111,184,255)",
		},
		{
			Color: c80color.Array{118,38,113,255},
			Expected:       "rgba(118,38,113,255)",
		},
		{
			Color: c80color.Array{44,181,233,255},
			Expected:       "rgba(44,181,233,255)",
		},
		{
			Color: c80color.Array{204,204,204,255},
			Expected:       "rgba(204,204,204,255)",
		},



		{
			Color: c80color.Array{255,255,255,255},
			Expected:       "rgba(255,255,255,255)",
		},
	}

	for testNumber, test := range tests {

		actual := test.Color.String()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
