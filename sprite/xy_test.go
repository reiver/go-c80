package c80sprite_test

import (
	"github.com/reiver/go-c80/sprite"

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
			ExpectedX: 0,
			ExpectedY: 1,
		},
		{
			Offset:    9,
			ExpectedX: 1,
			ExpectedY: 1,
		},
		{
			Offset:    10,
			ExpectedX: 2,
			ExpectedY: 1,
		},
		{
			Offset:    11,
			ExpectedX: 3,
			ExpectedY: 1,
		},
		{
			Offset:    12,
			ExpectedX: 4,
			ExpectedY: 1,
		},
		{
			Offset:    13,
			ExpectedX: 5,
			ExpectedY: 1,
		},
		{
			Offset:    14,
			ExpectedX: 6,
			ExpectedY: 1,
		},
		{
			Offset:    15,
			ExpectedX: 7,
			ExpectedY: 1,
		},



		{
			Offset:    16,
			ExpectedX: 0,
			ExpectedY: 2,
		},
		{
			Offset:    17,
			ExpectedX: 1,
			ExpectedY: 2,
		},
		{
			Offset:    18,
			ExpectedX: 2,
			ExpectedY: 2,
		},
		{
			Offset:    19,
			ExpectedX: 3,
			ExpectedY: 2,
		},
		{
			Offset:    20,
			ExpectedX: 4,
			ExpectedY: 2,
		},
		{
			Offset:    21,
			ExpectedX: 5,
			ExpectedY: 2,
		},
		{
			Offset:    22,
			ExpectedX: 6,
			ExpectedY: 2,
		},
		{
			Offset:    23,
			ExpectedX: 7,
			ExpectedY: 2,
		},



		{
			Offset:    24,
			ExpectedX: 0,
			ExpectedY: 3,
		},
		{
			Offset:    25,
			ExpectedX: 1,
			ExpectedY: 3,
		},
		{
			Offset:    26,
			ExpectedX: 2,
			ExpectedY: 3,
		},
		{
			Offset:    27,
			ExpectedX: 3,
			ExpectedY: 3,
		},
		{
			Offset:    28,
			ExpectedX: 4,
			ExpectedY: 3,
		},
		{
			Offset:    29,
			ExpectedX: 5,
			ExpectedY: 3,
		},
		{
			Offset:    30,
			ExpectedX: 6,
			ExpectedY: 3,
		},
		{
			Offset:    31,
			ExpectedX: 7,
			ExpectedY: 3,
		},



		{
			Offset:    32,
			ExpectedX: 0,
			ExpectedY: 4,
		},
		{
			Offset:    33,
			ExpectedX: 1,
			ExpectedY: 4,
		},
		{
			Offset:    34,
			ExpectedX: 2,
			ExpectedY: 4,
		},
		{
			Offset:    35,
			ExpectedX: 3,
			ExpectedY: 4,
		},
		{
			Offset:    36,
			ExpectedX: 4,
			ExpectedY: 4,
		},
		{
			Offset:    37,
			ExpectedX: 5,
			ExpectedY: 4,
		},
		{
			Offset:    38,
			ExpectedX: 6,
			ExpectedY: 4,
		},
		{
			Offset:    39,
			ExpectedX: 7,
			ExpectedY: 4,
		},



		{
			Offset:    40,
			ExpectedX: 0,
			ExpectedY: 5,
		},
		{
			Offset:    41,
			ExpectedX: 1,
			ExpectedY: 5,
		},
		{
			Offset:    42,
			ExpectedX: 2,
			ExpectedY: 5,
		},
		{
			Offset:    43,
			ExpectedX: 3,
			ExpectedY: 5,
		},
		{
			Offset:    44,
			ExpectedX: 4,
			ExpectedY: 5,
		},
		{
			Offset:    45,
			ExpectedX: 5,
			ExpectedY: 5,
		},
		{
			Offset:    46,
			ExpectedX: 6,
			ExpectedY: 5,
		},
		{
			Offset:    47,
			ExpectedX: 7,
			ExpectedY: 5,
		},


		{
			Offset:    48,
			ExpectedX: 0,
			ExpectedY: 6,
		},
		{
			Offset:    49,
			ExpectedX: 1,
			ExpectedY: 6,
		},
		{
			Offset:    50,
			ExpectedX: 2,
			ExpectedY: 6,
		},
		{
			Offset:    51,
			ExpectedX: 3,
			ExpectedY: 6,
		},
		{
			Offset:    52,
			ExpectedX: 4,
			ExpectedY: 6,
		},
		{
			Offset:    53,
			ExpectedX: 5,
			ExpectedY: 6,
		},
		{
			Offset:    54,
			ExpectedX: 6,
			ExpectedY: 6,
		},
		{
			Offset:    55,
			ExpectedX: 7,
			ExpectedY: 6,
		},


		{
			Offset:    56,
			ExpectedX: 0,
			ExpectedY: 7,
		},
		{
			Offset:    57,
			ExpectedX: 1,
			ExpectedY: 7,
		},
		{
			Offset:    58,
			ExpectedX: 2,
			ExpectedY: 7,
		},
		{
			Offset:    59,
			ExpectedX: 3,
			ExpectedY: 7,
		},
		{
			Offset:    60,
			ExpectedX: 4,
			ExpectedY: 7,
		},
		{
			Offset:    61,
			ExpectedX: 5,
			ExpectedY: 7,
		},
		{
			Offset:    62,
			ExpectedX: 6,
			ExpectedY: 7,
		},
		{
			Offset:    63,
			ExpectedX: 7,
			ExpectedY: 7,
		},

	}

	for testNumber, test := range tests {

		actualX, actualY := c80sprite.XY(test.Offset)

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
