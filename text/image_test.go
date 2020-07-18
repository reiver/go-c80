package c80text_test

import (
	"github.com/reiver/go-c80/text"

	"image"

	"testing"
)

func TestImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS>
	var x image.Image = c80text.Type{}

	if nil == x {
		t.Error("This should never happen.")
		return
	}
}
