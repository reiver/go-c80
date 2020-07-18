package c80pixel_test

import (
	"github.com/reiver/go-c80/pixel"

	"image/color"

	"testing"
)

func TestColor(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS>
	var x color.Color = c80pixel.Type{}

	if nil == x {
		t.Error("This should never happen.")
		return
	}
}
