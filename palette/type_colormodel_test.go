package c80palette_test

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/palette"

	"math/rand"
	"time"

	"testing"
)

func TestTypeColorModel(t *testing.T) {
	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var buffer [c80palette.Size*c80color.ByteSize]uint8
	for i:=0; i<len(buffer); i++ {
		buffer[i] = uint8(randomness.Intn(256))
	}

	var palette c80palette.Type
	var err error

	palette, err = c80palette.Wrap(buffer[:])
	if nil != err {
		t.Errorf("Received an error, but not actually expect one.")
		t.Logf("ERROR: (%T) %q", err, err)
		return
	}

	colorModel := palette.ColorModel()

	for index:=0; index<c80palette.Size; index++ {
		color := palette.Color(uint8(index))

		er, eg, eb, ea := color.RGBA()

		actualColor := colorModel.Convert(color)

		ar, ag, ab, aa := actualColor.RGBA()

		if er != ar || eg != ag || eb != ab || ea != aa {
			t.Errorf("For index=%d, actual color was not what was expected.", index)
			t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", er, eg, eb, ea)
			t.Logf("ACTUAL:   (r,g,b,a)=(%d,%d,%d,%d)", ar, ag, ab, aa)
			continue
		}
	}
}
