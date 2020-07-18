package c80pixel_test

import (
	"github.com/reiver/go-c80/pixel"

	"image"

	"testing"
)

func TestImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS>
	var x image.Image = c80pixel.Type{}

	if nil == x {
		t.Error("This should never happen.")
		return
	}
}
