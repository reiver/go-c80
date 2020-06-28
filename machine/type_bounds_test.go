package c80machine_test

import (
	"github.com/reiver/go-c80/machine"

	"image"

	"testing"
)

func TestTypeBounds(t *testing.T) {

	expected := image.Rectangle{
		Min: image.Point{X:0       , Y:0       },
		Max: image.Point{X:(128-1) , Y:(192-1) },
	}

	var machine c80machine.Type

	actual := machine.Bounds()

	if expected != actual {
		t.Errorf("Actual bounds was not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
		return
	}
}
