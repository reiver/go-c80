package c80sheet32x32

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
			Expected: 1024,
		},
		{
			N:        2,
			Expected: 2*1024,
		},
		{
			N:        3,
			Expected: 3*1024,
		},
		{
			N:        4,
			Expected: 4*1024,
		},
		{
			N:        5,
			Expected: 5*1024,
		},
		{
			N:        6,
			Expected: 6*1024,
		},
		{
			N:        7,
			Expected: 7*1024,
		},



		{
			N:        8,
			Expected: 8*1024,
		},
		{
			N:        9,
			Expected: 9*1024,
		},
		{
			N:        10,
			Expected: 10*1024,
		},
		{
			N:        11,
			Expected: 11*1024,
		},
		{
			N:        12,
			Expected: 12*1024,
		},
		{
			N:        13,
			Expected: 13*1024,
		},
		{
			N:        14,
			Expected: 14*1024,
		},
		{
			N:        15,
			Expected: 15*1024,
		},



		{
			N:        16,
			Expected: 16*1024,
		},
		{
			N:        17,
			Expected: 17*1024,
		},
		{
			N:        18,
			Expected: 18*1024,
		},
		{
			N:        19,
			Expected: 19*1024,
		},
		{
			N:        20,
			Expected: 20*1024,
		},
		{
			N:        21,
			Expected: 21*1024,
		},
		{
			N:        22,
			Expected: 22*1024,
		},
		{
			N:        23,
			Expected: 23*1024,
		},






		{
			N:        240,
			Expected: 240*1024,
		},
		{
			N:        241,
			Expected: 241*1024,
		},
		{
			N:        242,
			Expected: 242*1024,
		},
		{
			N:        243,
			Expected: 243*1024,
		},
		{
			N:        244,
			Expected: 244*1024,
		},
		{
			N:        245,
			Expected: 245*1024,
		},
		{
			N:        246,
			Expected: 246*1024,
		},
		{
			N:        247,
			Expected: 247*1024,
		},



		{
			N:        248,
			Expected: 248*1024,
		},
		{
			N:        249,
			Expected: 249*1024,
		},
		{
			N:        250,
			Expected: 250*1024,
		},
		{
			N:        251,
			Expected: 251*1024,
		},
		{
			N:        252,
			Expected: 252*1024,
		},
		{
			N:        253,
			Expected: 253*1024,
		},
		{
			N:        254,
			Expected: 254*1024,
		},
		{
			N:        255,
			Expected: 255*1024,
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
