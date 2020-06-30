package c80pixel_test

import (
	"github.com/reiver/go-c80/pixel"

	"testing"
)

func TestArray_Peek(t *testing.T) {

	tests := []struct{
		Value c80pixel.Array
		Expected uint8
	}{
		{
			Value: c80pixel.Array{0},
			Expected:             0,
		},
		{
			Value: c80pixel.Array{1},
			Expected:             1,
		},
		{
			Value: c80pixel.Array{2},
			Expected:             2,
		},
		{
			Value: c80pixel.Array{3},
			Expected:             3,
		},



		{
			Value: c80pixel.Array{4},
			Expected:             4,
		},
		{
			Value: c80pixel.Array{5},
			Expected:             5,
		},
		{
			Value: c80pixel.Array{6},
			Expected:             6,
		},
		{
			Value: c80pixel.Array{7},
			Expected:             7,
		},



		{
			Value: c80pixel.Array{8},
			Expected:             8,
		},
		{
			Value: c80pixel.Array{9},
			Expected:             9,
		},
		{
			Value: c80pixel.Array{10},
			Expected:             10,
		},
		{
			Value: c80pixel.Array{11},
			Expected:             11,
		},



		{
			Value: c80pixel.Array{12},
			Expected:             12,
		},
		{
			Value: c80pixel.Array{13},
			Expected:             13,
		},
		{
			Value: c80pixel.Array{14},
			Expected:             14,
		},
		{
			Value: c80pixel.Array{15},
			Expected:             15,
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.Peek()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual peeked value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
