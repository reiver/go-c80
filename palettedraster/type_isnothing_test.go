package c80palettedraster_test

import (
	"github.com/reiver/go-c80/palettedraster"

	"testing"
)

func TestType_Nothing(t *testing.T) {

	var palettedraster c80palettedraster.Type

	{
		expected := true
		actual   := palettedraster.IsNothing()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}

	{
		expected := false
		actual   := palettedraster.IsSomething()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}
}
