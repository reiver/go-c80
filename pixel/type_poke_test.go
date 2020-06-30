package c80pixel_test

import (
	"github.com/reiver/go-c80/pixel"

	"math/rand"
	"time"

	"testing"
)

func TestTypePoke(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<64; testNumber++ {

		var buffer c80pixel.Array = c80pixel.Array{
			uint8(randomness.Intn(16)),
		}

		var pixel  c80pixel.Type = c80pixel.Type(buffer[:])

		expected := uint8(randomness.Intn(16))

		pixel.Poke(expected)

		actual := pixel.Peek()

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
