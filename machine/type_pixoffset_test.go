package c80machine_test

import (
	"github.com/reiver/go-c80/machine"

	"testing"
)

func TestTypePixOffset(t *testing.T) {

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
			X:        120,
			Y:          0,
			Expected: 120,
		},
		{
			X:        121,
			Y:          0,
			Expected: 121,
		},
		{
			X:        122,
			Y:          0,
			Expected: 122,
		},
		{
			X:        123,
			Y:          0,
			Expected: 123,
		},
		{
			X:        124,
			Y:          0,
			Expected: 124,
		},
		{
			X:        125,
			Y:          0,
			Expected: 125,
		},
		{
			X:        126,
			Y:          0,
			Expected: 126,
		},
		{
			X:        127,
			Y:          0,
			Expected: 127,
		},



		{
			X:          0,
			Y:          1,
			Expected: 128,
		},
		{
			X:          1,
			Y:          1,
			Expected: 129,
		},
		{
			X:          2,
			Y:          1,
			Expected: 130,
		},
		{
			X:          3,
			Y:          1,
			Expected: 131,
		},
		{
			X:          4,
			Y:          1,
			Expected: 132,
		},
		{
			X:          5,
			Y:          1,
			Expected: 133,
		},



		{
			X:          0,
			Y:          2,
			Expected: 256,
		},



		{
			X:          0,
			Y:          3,
			Expected: 384,
		},



		{
			X:          0,
			Y:          4,
			Expected: 512,
		},
		{
			X:          1,
			Y:          4,
			Expected: 513,
		},
		{
			X:          2,
			Y:          4,
			Expected: 514,
		},
		{
			X:          3,
			Y:          4,
			Expected: 515,
		},
		{
			X:          4,
			Y:          4,
			Expected: 516,
		},
		{
			X:          5,
			Y:          4,
			Expected: 517,
		},



		{
			X:          120,
			Y:          191,
			Expected:  (128*192-8),
		},
		{
			X:          121,
			Y:          191,
			Expected:  (128*192-7),
		},
		{
			X:          122,
			Y:          191,
			Expected:  (128*192-6),
		},
		{
			X:          123,
			Y:          191,
			Expected:  (128*192-5),
		},
		{
			X:          124,
			Y:          191,
			Expected:  (128*192-4),
		},
		{
			X:          125,
			Y:          191,
			Expected:  (128*192-3),
		},
		{
			X:          126,
			Y:          191,
			Expected:  (128*192-2),
		},
		{
			X:          127,
			Y:          191,
			Expected:  (128*192-1),
		},
	}

	for testNumber, test := range tests {

		var machine c80machine.Type

		actual := machine.PixOffset(test.X, test.Y)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, actual pix offset was not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
