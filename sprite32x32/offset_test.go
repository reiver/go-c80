package c80sprite32x32

import (
	"testing"
)

func TestOffset(t *testing.T) {

	tests := []struct{
		X int
		Y int
		Expected int
	}{
		{
			X:         0,
			Y:         0,
			Expected:  0,
		},
		{
			X:         1,
			Y:         0,
			Expected:  1,
		},
		{
			X:         2,
			Y:         0,
			Expected:  2,
		},
		{
			X:         3,
			Y:         0,
			Expected:  3,
		},
		{
			X:         4,
			Y:         0,
			Expected:  4,
		},
		{
			X:         5,
			Y:         0,
			Expected:  5,
		},
		{
			X:         6,
			Y:         0,
			Expected:  6,
		},
		{
			X:         7,
			Y:         0,
			Expected:  7,
		},



		{
			X:         8,
			Y:         0,
			Expected:  8,
		},
		{
			X:         9,
			Y:         0,
			Expected:  9,
		},
		{
			X:         10,
			Y:         0,
			Expected:  10,
		},
		{
			X:         11,
			Y:         0,
			Expected:  11,
		},
		{
			X:         12,
			Y:         0,
			Expected:  12,
		},
		{
			X:         13,
			Y:         0,
			Expected:  13,
		},
		{
			X:         14,
			Y:         0,
			Expected:  14,
		},
		{
			X:         15,
			Y:         0,
			Expected:  15,
		},



		{
			X:         16,
			Y:         0,
			Expected:  16,
		},
		{
			X:         17,
			Y:         0,
			Expected:  17,
		},
		{
			X:         18,
			Y:         0,
			Expected:  18,
		},
		{
			X:         19,
			Y:         0,
			Expected:  19,
		},
		{
			X:         20,
			Y:         0,
			Expected:  20,
		},
		{
			X:         21,
			Y:         0,
			Expected:  21,
		},
		{
			X:         22,
			Y:         0,
			Expected:  22,
		},
		{
			X:         23,
			Y:         0,
			Expected:  23,
		},



		{
			X:         24,
			Y:         0,
			Expected:  24,
		},
		{
			X:         25,
			Y:         0,
			Expected:  25,
		},
		{
			X:         26,
			Y:         0,
			Expected:  26,
		},
		{
			X:         27,
			Y:         0,
			Expected:  27,
		},
		{
			X:         28,
			Y:         0,
			Expected:  28,
		},
		{
			X:         29,
			Y:         0,
			Expected:  29,
		},
		{
			X:         30,
			Y:         0,
			Expected:  30,
		},
		{
			X:         31,
			Y:         0,
			Expected:  31,
		},



		{
			X:         0,
			Y:         1,
			Expected:  32,
		},
		{
			X:         1,
			Y:         1,
			Expected:  33,
		},
		{
			X:         2,
			Y:         1,
			Expected:  34,
		},
		{
			X:         3,
			Y:         1,
			Expected:  35,
		},
		{
			X:         4,
			Y:         1,
			Expected:  36,
		},
		{
			X:         5,
			Y:         1,
			Expected:  37,
		},
		{
			X:         6,
			Y:         1,
			Expected:  38,
		},
		{
			X:         7,
			Y:         1,
			Expected:  39,
		},



		{
			X:         31,
			Y:         31,
			Expected:  1023,
		},
	}

	for testNumber, test := range tests {

		actual := offset(test.X, test.Y)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, actual value was not what was expected.", testNumber)
			t.Logf("(x,y) = (%d, %d)", test.X, test.Y)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
