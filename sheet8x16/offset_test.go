package c80sheet8x16

import (
	"testing"
)

func TestOffset(t *testing.T) {

	tests := []struct{
		N uint8
		Expected int
	}{
		{
			N:        0,
			Expected: 0,
		},
		{
			N:        1,
			Expected: 128,
		},
		{
			N:        2,
			Expected: 2*128,
		},
		{
			N:        3,
			Expected: 3*128,
		},
		{
			N:        4,
			Expected: 4*128,
		},
		{
			N:        5,
			Expected: 5*128,
		},
		{
			N:        6,
			Expected: 6*128,
		},
		{
			N:        7,
			Expected: 7*128,
		},



		{
			N:        8,
			Expected: 8*128,
		},
		{
			N:        9,
			Expected: 9*128,
		},
		{
			N:        10,
			Expected: 10*128,
		},
		{
			N:        11,
			Expected: 11*128,
		},
		{
			N:        12,
			Expected: 12*128,
		},
		{
			N:        13,
			Expected: 13*128,
		},
		{
			N:        14,
			Expected: 14*128,
		},
		{
			N:        15,
			Expected: 15*128,
		},



		{
			N:        16,
			Expected: 16*128,
		},
		{
			N:        17,
			Expected: 17*128,
		},
		{
			N:        18,
			Expected: 18*128,
		},
		{
			N:        19,
			Expected: 19*128,
		},
		{
			N:        20,
			Expected: 20*128,
		},
		{
			N:        21,
			Expected: 21*128,
		},
		{
			N:        22,
			Expected: 22*128,
		},
		{
			N:        23,
			Expected: 23*128,
		},






		{
			N:        240,
			Expected: 240*128,
		},
		{
			N:        241,
			Expected: 241*128,
		},
		{
			N:        242,
			Expected: 242*128,
		},
		{
			N:        243,
			Expected: 243*128,
		},
		{
			N:        244,
			Expected: 244*128,
		},
		{
			N:        245,
			Expected: 245*128,
		},
		{
			N:        246,
			Expected: 246*128,
		},
		{
			N:        247,
			Expected: 247*128,
		},



		{
			N:        248,
			Expected: 248*128,
		},
		{
			N:        249,
			Expected: 249*128,
		},
		{
			N:        250,
			Expected: 250*128,
		},
		{
			N:        251,
			Expected: 251*128,
		},
		{
			N:        252,
			Expected: 252*128,
		},
		{
			N:        253,
			Expected: 253*128,
		},
		{
			N:        254,
			Expected: 254*128,
		},
		{
			N:        255,
			Expected: 255*128,
		},
	}

	for testNumber, test := range tests {
		actual := offset(test.N)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value was not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
