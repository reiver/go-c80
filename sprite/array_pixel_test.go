package c80sprite_test

import (
	"github.com/reiver/go-c80/sprite"

	"math/rand"
	"time"

	"testing"
)

func TestArray_Pixel(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var sprite c80sprite.Array

	for i:=0; i<len(sprite); i++ {
		sprite[i] = uint8(randomness.Intn(16))
	}

	{
		numNotZero := 0

		for y:=0; y<c80sprite.Height; y++ {
			for x:=0; x<c80sprite.Width; x++ {
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

	for y:=0; y<c80sprite.Height; y++ {
		for x:=0; x<c80sprite.Width; x++ {

			expectedIndex := uint8(randomness.Intn(16))

			sprite.Pixel(x,y).Poke(expectedIndex)

			actualIndex := sprite.Pixel(x,y).Peek()

			if expected, actual := expectedIndex, actualIndex; expected != actual {
				t.Errorf("For pixel (x,y)=(%d,%d) in sprite, actual index is not what was expected.", x, y)
				t.Logf("EXPECTED: %d, ", expected)
				t.Logf("ACTUAL:   %d, ", actual)
				continue
			}
		}
	}
}
