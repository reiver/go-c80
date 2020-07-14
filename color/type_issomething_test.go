package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"testing"
)

func TestType_IsSomething(t *testing.T) {

	var buffer [4]uint8 = [4]uint8{11,22,33,44}

	var color c80color.Type
	var err error

	color, err = c80color.Wrap(buffer[:])
	if nil != err {
		t.Errorf("Received an error, but did not actually expect one/")
		return
	}

	{
		expected := true
		actual   := color.IsSomething()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}

	{
		expected := false
		actual   := color.IsNothing()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}
}
