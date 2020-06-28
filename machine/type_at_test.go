package c80machine_test

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/machine"
	"github.com/reiver/go-c80/palette"

	"math/rand"
	"time"

	"testing"
)

func TestTypeAt(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var palette [c80palette.Size]c80color.Type
	for index:=0; index<c80palette.Size; index++ {
		palette[index] = c80color.RGB(
			uint8(randomness.Intn(256)),
			uint8(randomness.Intn(256)),
			uint8(randomness.Intn(256)),
		)
	}

	for index:=uint8(0); index<c80palette.Size; index++ {

		testNumber := index

		var machine c80machine.Type
		machine.SetPalette(c80palette.RGBs(palette[:]...))
		machine.Dye(index)

		for y:=0; y<c80machine.Height; y++ {
			for x:=0; x<c80machine.Width; x++ {

				color := machine.At(x,y)

				ar,ag,ab,aa := color.RGBA()

				er,eg,eb,ea := palette[index].RGBA()

				if er != ar || eg != ag || eb != ab || ea != aa {
					t.Errorf("For test #%d, actual color for pixel (x,y)=(%d,%d) was not what was expected.", testNumber, x,y)
					t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", er,eg,eb,ea)
					t.Logf("ACTUAL:   (r,g,b,a)=(%d,%d,%d,%d)", ar,ag,ab,aa)
				}
			}
		}
	}

}