package c80raster_test

import (
	"github.com/reiver/go-c80/raster"

	"testing"
)

func TestTypeString(t *testing.T) {

	tests := []struct{
		Raster c80raster.Type
		Expected string
	}{
		{
			Raster: c80raster.Type{},
			Expected: "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAH8AAAC/BAMAAAAlex6UAAAAMFBMVEUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABaPxwLAAAAJ0lEQVR4nOzAgQAAAADCsPypMzjBNgAAAAAAAAAAAAAAgPYAAAD//zB/AAGXmfMxAAAAAElFTkSuQmCC",
		},
	}

	for testNumber, test := range tests {

		actual := test.Raster.String()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value was not what was expected." , testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
