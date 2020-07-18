package c80sprite8x8_test

import (
	"github.com/reiver/go-c80/sprite8x8"

	"math/rand"
	"time"

	"testing"
)

func TestType_Pel(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var buffer [64]uint8
	var sprite c80sprite8x8.Type
	var err error

	sprite, err = c80sprite8x8.Wrap(buffer[:])
	if nil != err {
		t.Errorf("Received an error, but did not actually expect one.")
		t.Logf("ERROR: (%T) %q", err, err)
		return
	}

	for i:=0; i<len(buffer); i++ {
		buffer[i] = uint8(randomness.Intn(16))
	}

	{
		numNotZero := 0

		for y:=0; y<c80sprite8x8.Height; y++ {
			for x:=0; x<c80sprite8x8.Width; x++ {
				index := sprite.Pel(x,y).Peek()

				if 0 != index {
					numNotZero++
				}
			}
		}

		if 0 == numNotZero {
			t.Errorf("Probably peeking on the pel is not working... probably.")
			t.Logf("Number of non-zero color palette indexes: %d", numNotZero)
			return
		}
	}

	for y:=0; y<c80sprite8x8.Height; y++ {
		for x:=0; x<c80sprite8x8.Width; x++ {

			expectedIndex := uint8(randomness.Intn(16))

			sprite.Pel(x,y).Poke(expectedIndex)

			actualIndex := sprite.Pel(x,y).Peek()

			if expected, actual := expectedIndex, actualIndex; expected != actual {
				t.Errorf("For pel (x,y)=(%d,%d) in sprite, actual index is not what was expected.", x, y)
				t.Logf("EXPECTED: %d, ", expected)
				t.Logf("ACTUAL:   %d, ", actual)
				continue
			}
		}
	}
}
