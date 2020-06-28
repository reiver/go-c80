package c80machine_test

import (
	"github.com/reiver/go-c80/machine"

	"math/rand"
	"time"

	"testing"
)

func TestTypeSetColorIndex(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var machine c80machine.Type

	for y := 0; y<192; y++ {
		for x := 0; x<128; x++  {

			{
				actual := machine.ColorIndexAt(x,y)

				if expected := uint8(0); expected != actual {
					t.Errorf("For test at (x,y)=(%d,%d) actual initial color index not what was expected.", x, y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}

			{
				index := uint8(randomness.Intn(16))

				machine.SetColorIndex(x,y, index)

				actual := machine.ColorIndexAt(x,y)

				if expected := index; expected != actual {
					t.Errorf("For test at (x,y)=(%d,%d) actual color index not what was expected.", x, y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}

			{
				index := uint8(randomness.Intn(16))

				machine.SetColorIndex(x,y, index)

				actual := machine.ColorIndexAt(x,y)

				if expected := index; expected != actual {
					t.Errorf("For test at (x,y)=(%d,%d) actual color index not what was expected.", x, y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}

			{
				index := uint8(2)

				machine.SetColorIndex(x,y, index)

				actual := machine.ColorIndexAt(x,y)

				if expected := index; expected != actual {
					t.Errorf("For test at (x,y)=(%d,%d) actual color index not what was expected.", x, y)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					continue
				}
			}
		}
	}
}
