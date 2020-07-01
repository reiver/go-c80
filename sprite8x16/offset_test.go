package c80sprite8x16_test

import (
	"github.com/reiver/go-c80/sprite8x16"

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
			X:         0,
			Y:         1,
			Expected:  8,
		},
		{
			X:         1,
			Y:         1,
			Expected:  9,
		},
		{
			X:         2,
			Y:         1,
			Expected:  10,
		},
		{
			X:         3,
			Y:         1,
			Expected:  11,
		},
		{
			X:         4,
			Y:         1,
			Expected:  12,
		},
		{
			X:         5,
			Y:         1,
			Expected:  13,
		},
		{
			X:         6,
			Y:         1,
			Expected:  14,
		},
		{
			X:         7,
			Y:         1,
			Expected:  15,
		},



		{
			X:         0,
			Y:         2,
			Expected:  16,
		},
		{
			X:         1,
			Y:         2,
			Expected:  17,
		},
		{
			X:         2,
			Y:         2,
			Expected:  18,
		},
		{
			X:         3,
			Y:         2,
			Expected:  19,
		},
		{
			X:         4,
			Y:         2,
			Expected:  20,
		},
		{
			X:         5,
			Y:         2,
			Expected:  21,
		},
		{
			X:         6,
			Y:         2,
			Expected:  22,
		},
		{
			X:         7,
			Y:         2,
			Expected:  23,
		},



		{
			X:         0,
			Y:         3,
			Expected:  24,
		},
		{
			X:         1,
			Y:         3,
			Expected:  25,
		},
		{
			X:         2,
			Y:         3,
			Expected:  26,
		},
		{
			X:         3,
			Y:         3,
			Expected:  27,
		},
		{
			X:         4,
			Y:         3,
			Expected:  28,
		},
		{
			X:         5,
			Y:         3,
			Expected:  29,
		},
		{
			X:         6,
			Y:         3,
			Expected:  30,
		},
		{
			X:         7,
			Y:         3,
			Expected:  31,
		},



		{
			X:         0,
			Y:         4,
			Expected:  32,
		},
		{
			X:         1,
			Y:         4,
			Expected:  33,
		},
		{
			X:         2,
			Y:         4,
			Expected:  34,
		},
		{
			X:         3,
			Y:         4,
			Expected:  35,
		},
		{
			X:         4,
			Y:         4,
			Expected:  36,
		},
		{
			X:         5,
			Y:         4,
			Expected:  37,
		},
		{
			X:         6,
			Y:         4,
			Expected:  38,
		},
		{
			X:         7,
			Y:         4,
			Expected:  39,
		},



		{
			X:         0,
			Y:         5,
			Expected:  40,
		},
		{
			X:         1,
			Y:         5,
			Expected:  41,
		},
		{
			X:         2,
			Y:         5,
			Expected:  42,
		},
		{
			X:         3,
			Y:         5,
			Expected:  43,
		},
		{
			X:         4,
			Y:         5,
			Expected:  44,
		},
		{
			X:         5,
			Y:         5,
			Expected:  45,
		},
		{
			X:         6,
			Y:         5,
			Expected:  46,
		},
		{
			X:         7,
			Y:         5,
			Expected:  47,
		},



		{
			X:         0,
			Y:         6,
			Expected:  48,
		},
		{
			X:         1,
			Y:         6,
			Expected:  49,
		},
		{
			X:         2,
			Y:         6,
			Expected:  50,
		},
		{
			X:         3,
			Y:         6,
			Expected:  51,
		},
		{
			X:         4,
			Y:         6,
			Expected:  52,
		},
		{
			X:         5,
			Y:         6,
			Expected:  53,
		},
		{
			X:         6,
			Y:         6,
			Expected:  54,
		},
		{
			X:         7,
			Y:         6,
			Expected:  55,
		},



		{
			X:         0,
			Y:         7,
			Expected:  56,
		},
		{
			X:         1,
			Y:         7,
			Expected:  57,
		},
		{
			X:         2,
			Y:         7,
			Expected:  58,
		},
		{
			X:         3,
			Y:         7,
			Expected:  59,
		},
		{
			X:         4,
			Y:         7,
			Expected:  60,
		},
		{
			X:         5,
			Y:         7,
			Expected:  61,
		},
		{
			X:         6,
			Y:         7,
			Expected:  62,
		},
		{
			X:         7,
			Y:         7,
			Expected:  63,
		},
	}

	for testNumber, test := range tests {

		actual := c80sprite8x16.Offset(test.X, test.Y)

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
