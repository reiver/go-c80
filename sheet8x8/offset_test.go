package c80sheet8x8

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
			Expected: 64,
		},
		{
			N:        2,
			Expected: 2*64,
		},
		{
			N:        3,
			Expected: 3*64,
		},
		{
			N:        4,
			Expected: 4*64,
		},
		{
			N:        5,
			Expected: 5*64,
		},
		{
			N:        6,
			Expected: 6*64,
		},
		{
			N:        7,
			Expected: 7*64,
		},



		{
			N:        8,
			Expected: 8*64,
		},
		{
			N:        9,
			Expected: 9*64,
		},
		{
			N:        10,
			Expected: 10*64,
		},
		{
			N:        11,
			Expected: 11*64,
		},
		{
			N:        12,
			Expected: 12*64,
		},
		{
			N:        13,
			Expected: 13*64,
		},
		{
			N:        14,
			Expected: 14*64,
		},
		{
			N:        15,
			Expected: 15*64,
		},



		{
			N:        16,
			Expected: 16*64,
		},
		{
			N:        17,
			Expected: 17*64,
		},
		{
			N:        18,
			Expected: 18*64,
		},
		{
			N:        19,
			Expected: 19*64,
		},
		{
			N:        20,
			Expected: 20*64,
		},
		{
			N:        21,
			Expected: 21*64,
		},
		{
			N:        22,
			Expected: 22*64,
		},
		{
			N:        23,
			Expected: 23*64,
		},






		{
			N:        240,
			Expected: 240*64,
		},
		{
			N:        241,
			Expected: 241*64,
		},
		{
			N:        242,
			Expected: 242*64,
		},
		{
			N:        243,
			Expected: 243*64,
		},
		{
			N:        244,
			Expected: 244*64,
		},
		{
			N:        245,
			Expected: 245*64,
		},
		{
			N:        246,
			Expected: 246*64,
		},
		{
			N:        247,
			Expected: 247*64,
		},



		{
			N:        248,
			Expected: 248*64,
		},
		{
			N:        249,
			Expected: 249*64,
		},
		{
			N:        250,
			Expected: 250*64,
		},
		{
			N:        251,
			Expected: 251*64,
		},
		{
			N:        252,
			Expected: 252*64,
		},
		{
			N:        253,
			Expected: 253*64,
		},
		{
			N:        254,
			Expected: 254*64,
		},
		{
			N:        255,
			Expected: 255*64,
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
