package c80raster_test

import (
	"github.com/reiver/go-c80/raster"

	"image"

	"testing"
)

func TestTypeBounds(t *testing.T) {

	expected := image.Rectangle{
		Min: image.Point{X:0       , Y:0       },
		Max: image.Point{X:(128-1) , Y:(192-1) },
	}

	var raster c80raster.Type

	actual := raster.Bounds()

	if expected != actual {
		t.Errorf("Actual bounds was not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
		return
	}
}
