package c80machine_test

import (
	"github.com/reiver/go-c80/machine"

	"testing"
)

func TestTypeString(t *testing.T) {

	tests := []struct{
		Machine c80machine.Type
		Expected string
	}{
		{
			Machine: c80machine.Type{},
			Expected: "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAH8AAAC/BAMAAAAlex6UAAAAMFBMVEUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABaPxwLAAAAJ0lEQVR4nOzAgQAAAADCsPypMzjBNgAAAAAAAAAAAAAAgPYAAAD//zB/AAGXmfMxAAAAAElFTkSuQmCC",
		},
	}

	for testNumber, test := range tests {

		actual := test.Machine.String()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value was not what was expected." , testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
