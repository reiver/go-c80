package c80machine_test

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/machine"
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"

	"math/rand"
	"time"

	"testing"
)

func TestTypeAt(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var machine c80machine.Type

	for index:=uint8(0); index<c80palette.Size; index++ {
		machine.Palette().Color(index).Poke(
			uint8(randomness.Intn(256)),
			uint8(randomness.Intn(256)),
			uint8(randomness.Intn(256)),
			uint8(255),
		)
	}

	for index:=uint8(0); index<c80palette.Size; index++ {
		actual := machine.Palette().Color(index).Peek()
		notExpected := c80color.Array{}

		if notExpected == actual {
			t.Errorf("Was not able to set color №%d in palette.", index)
			t.Log("NOT EXPECTED:", notExpected)
			t.Log("ACTUAL:      ", actual)
			continue
		}
	}

	for y:=0; y<c80raster.Height; y++ {
		for x:=0; x<c80raster.Width; x++ {

			index := uint8(randomness.Intn(16))

			machine.Raster0().SetColorIndex(x, y, index)
		}
	}

	{
		numNotZero := 0

		for y:=0; y<c80raster.Height; y++ {
			for x:=0; x<c80raster.Width; x++ {

				index := machine.Raster0().ColorIndexAt(x,y)

				if 0 != index {
					numNotZero++
				}
			}
		}

		if 0 == numNotZero {
			t.Errorf("Was not able to set color of pixels on raster №0.")
			t.Log("Number of non-zero pixels:", numNotZero)
			return
		}
	}

	for testNumber:=0; testNumber<50; testNumber++ {
		x := randomness.Intn(c80raster.Width)
		y := randomness.Intn(c80raster.Height)

		index := uint8(randomness.Intn(c80palette.Size))
		color := machine.Palette().Color(index)

		er,eg,eb,ea := color.RGBA()

		machine.Raster0().SetColorIndex(x,y,index)

		ar,ag,ab,aa := machine.At(x,y).RGBA()

		if er != ar || eg != ag || eb != ab || ea != aa {
			t.Errorf("For test #%d, actual RGBA is not what was actually expected.", testNumber)
			t.Logf("EXPECTED: RGBA16(%d,%d,%d,%d)", er,eg,eb,ea)
			t.Logf("ACTUAL:   RGBA16(%d,%d,%d,%d)", ar,ag,ab,aa)
			continue
		}
	}
}

