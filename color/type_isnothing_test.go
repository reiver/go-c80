package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"testing"
)

func TestType_Nothing(t *testing.T) {

	var color c80color.Type

	{
		expected := true
		actual   := color.IsNothing()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}

	{
		expected := false
		actual   := color.IsSomething()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}
}
