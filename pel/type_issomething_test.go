package c80pel_test

import (
	"github.com/reiver/go-c80/pel"

	"testing"
)

func TestType_IsSomething(t *testing.T) {

	var buffer [1]uint8 = [1]uint8{5}

	var pel c80pel.Type
	var err error

	pel, err = c80pel.Wrap(buffer[:])
	if nil != err {
		t.Errorf("Received an error, but did not actually expect one.")
		t.Logf("ERROR: (%T) %q", err, err)
		return
	}

	{
		expected := true
		actual   := pel.IsSomething()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}

	{
		expected := false
		actual   := pel.IsNothing()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}
}
