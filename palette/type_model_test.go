package c80palette_test

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/palette"

	"math/rand"
	"time"

	"testing"
)

func TestTypeModel(t *testing.T) {
	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var colors [c80palette.Size]c80color.Type
	for index:=0; index<c80palette.Size; index++ {
		colors[index] = c80color.RGB(
			uint8(randomness.Intn(256)),
			uint8(randomness.Intn(256)),
			uint8(randomness.Intn(256)),
		)
	}

	var palette c80palette.Type = c80palette.RGBs(colors[:]...)

	model := palette.Model()


	for index, color := range colors {
		er, eg, eb, ea := color.RGBA()

		actualColor := model.Convert(color)

		ar, ag, ab, aa := actualColor.RGBA()

		if er != ar || eg != ag || eb != ab || ea != aa {
			t.Errorf("For index=%d, actual color was not what was expected.", index)
			t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", er, eg, eb, ea)
			t.Logf("ACTUAL:   (r,g,b,a)=(%d,%d,%d,%d)", ar, ag, ab, aa)
			continue
		}
	}
}
