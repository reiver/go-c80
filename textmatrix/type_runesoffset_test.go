package c80textmatrix

import (
	"testing"
)

func TestType_runesOffset(t *testing.T) {

	tests := []struct{
		X int
		Y int
		Expected int
	}{
		{
			X:        0,
			Y:        0,
			Expected: 0,
		},
		{
			X:        1,
			Y:        0,
			Expected: 1,
		},
		{
			X:        2,
			Y:        0,
			Expected: 2,
		},
		{
			X:        3,
			Y:        0,
			Expected: 3,
		},
		{
			X:        4,
			Y:        0,
			Expected: 4,
		},
		{
			X:        5,
			Y:        0,
			Expected: 5,
		},



		{
			X:        29,
			Y:        0,
			Expected: 29,
		},
		{
			X:        30,
			Y:        0,
			Expected: 30,
		},
		{
			X:        31,
			Y:        0,
			Expected: 31,
		},
		{
			X:        0,
			Y:        1,
			Expected: 32,
		},
		{
			X:        1,
			Y:        1,
			Expected: 33,
		},
		{
			X:        2,
			Y:        1,
			Expected: 34,
		},
		{
			X:        3,
			Y:        1,
			Expected: 35,
		},
		{
			X:        4,
			Y:        1,
			Expected: 36,
		},
		{
			X:        5,
			Y:        1,
			Expected: 37,
		},



		{
			X:        0,
			Y:        2,
			Expected: 64,
		},



		{
			X:        0,
			Y:        3,
			Expected: 96,
		},



		{
			X:        0,
			Y:        4,
			Expected: 128,
		},



		{
			X:        0,
			Y:        5,
			Expected: 160,
		},



		{
			X:        0,
			Y:        6,
			Expected: 192,
		},



		{
			X:        0,
			Y:        7,
			Expected: 224,
		},



		{
			X:        0,
			Y:        8,
			Expected: 256,
		},



		{
			X:        (Width       -5),
			Y:        (Height-1),
			Expected: (Width*Height-5),
		},
		{
			X:        (Width       -4),
			Y:        (Height-1),
			Expected: (Width*Height-4),
		},
		{
			X:        (Width       -3),
			Y:        (Height-1),
			Expected: (Width*Height-3),
		},
		{
			X:        (Width       -2),
			Y:        (Height-1),
			Expected: (Width*Height-2),
		},
		{
			X:        (Width       -1),
			Y:        (Height      -1),
			Expected: (Width*Height-1),
		},
	}

	for testNumber, test := range tests {

		var datum Type

		actual := datum.runesOffset(test.X,test.Y)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
			t.Log("(x,y) =", test.X, test.Y)
			t.Log("EXPECTED:", expected)
			t.Log("ACTUAL:  ", actual)
			continue
		}
	}
}
