package c80dye_test

import (
	"github.com/reiver/go-c80/dye"

	"image/color"

	"testing"
)

func TestColor(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS>
	var x color.Color = c80dye.Type{}

	if nil == x {
		t.Error("This should never happen.")
		return
	}
}
