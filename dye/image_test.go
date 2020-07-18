package c80dye_test

import (
	"github.com/reiver/go-c80/dye"

	"image"

	"testing"
)

func TestImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS>
	var x image.Image = c80dye.Type{}

	if nil == x {
		t.Error("This should never happen.")
		return
	}
}
