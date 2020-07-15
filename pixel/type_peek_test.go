package c80pixel_test

import (
	"github.com/reiver/go-c80/pixel"

	"testing"
)

func TestType_Peek(t *testing.T) {

	tests := []struct{
		Value [1]uint8
		Expected uint8
	}{
		{
			Value: [1]uint8{0},
			Expected:       0,
		},
		{
			Value: [1]uint8{1},
			Expected:       1,
		},
		{
			Value: [1]uint8{2},
			Expected:       2,
		},
		{
			Value: [1]uint8{3},
			Expected:       3,
		},



		{
			Value: [1]uint8{4},
			Expected:       4,
		},
		{
			Value: [1]uint8{5},
			Expected:       5,
		},
		{
			Value: [1]uint8{6},
			Expected:       6,
		},
		{
			Value: [1]uint8{7},
			Expected:       7,
		},



		{
			Value: [1]uint8{8},
			Expected:       8,
		},
		{
			Value: [1]uint8{9},
			Expected:       9,
		},
		{
			Value: [1]uint8{10},
			Expected:       10,
		},
		{
			Value: [1]uint8{11},
			Expected:       11,
		},



		{
			Value: [1]uint8{12},
			Expected:       12,
		},
		{
			Value: [1]uint8{13},
			Expected:       13,
		},
		{
			Value: [1]uint8{14},
			Expected:       14,
		},
		{
			Value: [1]uint8{15},
			Expected:       15,
		},
	}

	for testNumber, test := range tests {

		var pixel c80pixel.Type
		var err error

		pixel, err = c80pixel.Wrap(test.Value[:])
		if nil != err {
			t.Errorf("For test #%d, received an error, but did not actually expect one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}

		actual := pixel.Peek()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual peeked value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
