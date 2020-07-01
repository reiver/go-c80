package c80sprite32x32_test

import (
	"github.com/reiver/go-c80/sprite32x32"

	"testing"
)

func TestXY(t *testing.T) {

	tests := []struct{
		Offset int
		ExpectedX int
		ExpectedY int
	}{
		{
			Offset:    0,
			ExpectedX: 0,
			ExpectedY: 0,
		},
		{
			Offset:    1,
			ExpectedX: 1,
			ExpectedY: 0,
		},
		{
			Offset:    2,
			ExpectedX: 2,
			ExpectedY: 0,
		},
		{
			Offset:    3,
			ExpectedX: 3,
			ExpectedY: 0,
		},
		{
			Offset:    4,
			ExpectedX: 4,
			ExpectedY: 0,
		},
		{
			Offset:    5,
			ExpectedX: 5,
			ExpectedY: 0,
		},
		{
			Offset:    6,
			ExpectedX: 6,
			ExpectedY: 0,
		},
		{
			Offset:    7,
			ExpectedX: 7,
			ExpectedY: 0,
		},



		{
			Offset:    8,
			ExpectedX: 8,
			ExpectedY: 0,
		},
		{
			Offset:    9,
			ExpectedX: 9,
			ExpectedY: 0,
		},
		{
			Offset:    10,
			ExpectedX: 10,
			ExpectedY: 0,
		},
		{
			Offset:    11,
			ExpectedX: 11,
			ExpectedY: 0,
		},
		{
			Offset:    12,
			ExpectedX: 12,
			ExpectedY: 0,
		},
		{
			Offset:    13,
			ExpectedX: 13,
			ExpectedY: 0,
		},
		{
			Offset:    14,
			ExpectedX: 14,
			ExpectedY: 0,
		},
		{
			Offset:    15,
			ExpectedX: 15,
			ExpectedY: 0,
		},



		{
			Offset:    16,
			ExpectedX: 16,
			ExpectedY: 0,
		},
		{
			Offset:    17,
			ExpectedX: 17,
			ExpectedY: 0,
		},
		{
			Offset:    18,
			ExpectedX: 18,
			ExpectedY: 0,
		},
		{
			Offset:    19,
			ExpectedX: 19,
			ExpectedY: 0,
		},
		{
			Offset:    20,
			ExpectedX: 20,
			ExpectedY: 0,
		},
		{
			Offset:    21,
			ExpectedX: 21,
			ExpectedY: 0,
		},
		{
			Offset:    22,
			ExpectedX: 22,
			ExpectedY: 0,
		},
		{
			Offset:    23,
			ExpectedX: 23,
			ExpectedY: 0,
		},



		{
			Offset:    24,
			ExpectedX: 24,
			ExpectedY: 0,
		},
		{
			Offset:    25,
			ExpectedX: 25,
			ExpectedY: 0,
		},
		{
			Offset:    26,
			ExpectedX: 26,
			ExpectedY: 0,
		},
		{
			Offset:    27,
			ExpectedX: 27,
			ExpectedY: 0,
		},
		{
			Offset:    28,
			ExpectedX: 28,
			ExpectedY: 0,
		},
		{
			Offset:    29,
			ExpectedX: 29,
			ExpectedY: 0,
		},
		{
			Offset:    30,
			ExpectedX: 30,
			ExpectedY: 0,
		},
		{
			Offset:    31,
			ExpectedX: 31,
			ExpectedY: 0,
		},



		{
			Offset:    32,
			ExpectedX: 0,
			ExpectedY: 1,
		},
		{
			Offset:    33,
			ExpectedX: 1,
			ExpectedY: 1,
		},
		{
			Offset:    34,
			ExpectedX: 2,
			ExpectedY: 1,
		},
		{
			Offset:    35,
			ExpectedX: 3,
			ExpectedY: 1,
		},
		{
			Offset:    36,
			ExpectedX: 4,
			ExpectedY: 1,
		},
		{
			Offset:    37,
			ExpectedX: 5,
			ExpectedY: 1,
		},
		{
			Offset:    38,
			ExpectedX: 6,
			ExpectedY: 1,
		},
		{
			Offset:    39,
			ExpectedX: 7,
			ExpectedY: 1,
		},



		{
			Offset:    40,
			ExpectedX: 8,
			ExpectedY: 1,
		},
		{
			Offset:    41,
			ExpectedX: 9,
			ExpectedY: 1,
		},
		{
			Offset:    42,
			ExpectedX: 10,
			ExpectedY: 1,
		},
		{
			Offset:    43,
			ExpectedX: 11,
			ExpectedY: 1,
		},
		{
			Offset:    44,
			ExpectedX: 12,
			ExpectedY: 1,
		},
		{
			Offset:    45,
			ExpectedX: 13,
			ExpectedY: 1,
		},
		{
			Offset:    46,
			ExpectedX: 14,
			ExpectedY: 1,
		},
		{
			Offset:    47,
			ExpectedX: 15,
			ExpectedY: 1,
		},


		{
			Offset:    48,
			ExpectedX: 16,
			ExpectedY: 1,
		},
		{
			Offset:    49,
			ExpectedX: 17,
			ExpectedY: 1,
		},
		{
			Offset:    50,
			ExpectedX: 18,
			ExpectedY: 1,
		},
		{
			Offset:    51,
			ExpectedX: 19,
			ExpectedY: 1,
		},
		{
			Offset:    52,
			ExpectedX: 20,
			ExpectedY: 1,
		},
		{
			Offset:    53,
			ExpectedX: 21,
			ExpectedY: 1,
		},
		{
			Offset:    54,
			ExpectedX: 22,
			ExpectedY: 1,
		},
		{
			Offset:    55,
			ExpectedX: 23,
			ExpectedY: 1,
		},


		{
			Offset:    56,
			ExpectedX: 24,
			ExpectedY: 1,
		},
		{
			Offset:    57,
			ExpectedX: 25,
			ExpectedY: 1,
		},
		{
			Offset:    58,
			ExpectedX: 26,
			ExpectedY: 1,
		},
		{
			Offset:    59,
			ExpectedX: 27,
			ExpectedY: 1,
		},
		{
			Offset:    60,
			ExpectedX: 28,
			ExpectedY: 1,
		},
		{
			Offset:    61,
			ExpectedX: 29,
			ExpectedY: 1,
		},
		{
			Offset:    62,
			ExpectedX: 30,
			ExpectedY: 1,
		},
		{
			Offset:    63,
			ExpectedX: 31,
			ExpectedY: 1,
		},



		{
			Offset:    1016,
			ExpectedX: 24,
			ExpectedY: 31,
		},
		{
			Offset:    1017,
			ExpectedX: 25,
			ExpectedY: 31,
		},
		{
			Offset:    1018,
			ExpectedX: 26,
			ExpectedY: 31,
		},
		{
			Offset:    1019,
			ExpectedX: 27,
			ExpectedY: 31,
		},
		{
			Offset:    1020,
			ExpectedX: 28,
			ExpectedY: 31,
		},
		{
			Offset:    1021,
			ExpectedX: 29,
			ExpectedY: 31,
		},
		{
			Offset:    1022,
			ExpectedX: 30,
			ExpectedY: 31,
		},
		{
			Offset:    1023,
			ExpectedX: 31,
			ExpectedY: 31,
		},

	}

	for testNumber, test := range tests {

		actualX, actualY := c80sprite32x32.XY(test.Offset)

		expectedX := test.ExpectedX
		expectedY := test.ExpectedY

		if expectedX != actualX || expectedY != actualY {
			t.Errorf("For test #%d, actual value was not what was expected.", testNumber)
			t.Logf("OFFSET: %d", test.Offset)
			t.Logf("EXPECTED: (x,y)=(%d,%d)", expectedX, expectedY)
			t.Logf("ACTUAL:   (x,y)=(%d,%d)", actualX, actualY)
			continue
		}
	}
}
