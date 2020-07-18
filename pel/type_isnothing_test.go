package c80pel_test

import (
	"github.com/reiver/go-c80/pel"

	"testing"
)

func TestType_Nothing(t *testing.T) {

	var pel c80pel.Type

	{
		expected := true
		actual   := pel.IsNothing()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}

	{
		expected := false
		actual   := pel.IsSomething()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}
}
