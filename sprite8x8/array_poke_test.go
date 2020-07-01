package c80sprite8x8_test

import (
	"github.com/reiver/go-c80/sprite8x8"

	"math/rand"
	"time"

	"testing"
)

func TestArray_Poke(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var sprite c80sprite8x8.Array

	for i:=0; i<len(sprite); i++ {
		sprite[i] = uint8(randomness.Intn(16))
	}

	{
		numNotZero := 0

		for y:=0; y<c80sprite8x8.Height; y++ {
			for x:=0; x<c80sprite8x8.Width; x++ {
				index := sprite.Pixel(x,y).Peek()

				if 0 != index {
					numNotZero++
				}
			}
		}

		if 0 == numNotZero {
			t.Errorf("Probably peeking on the pixel is not working... probably.")
			t.Logf("Number of non-zero color palette indexes: %d", numNotZero)
			return
		}
	}

	{
		var expected [c80sprite8x8.Len]uint8
		for i:=0; i<len(expected); i++ {
			expected[i] = uint8(randomness.Intn(16))
		}

		sprite.Poke(expected[:]...)

		actual := [c80sprite8x8.Len]uint8(sprite)

		if expected != actual {
			t.Errorf("The actual value was not what was expected.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}

	}
}
