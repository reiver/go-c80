package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"math/rand"
	"time"

	"testing"
)

func TestTypePoke(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<50; testNumber++ {

		var buffer c80color.Array
		for i:=0; i<len(buffer); i++ {
			buffer[i] = uint8(randomness.Intn(256))
		}

		var expected c80color.Array
		for i:=0; i<len(expected); i++ {
			expected[i] = uint8(randomness.Intn(256))
		}

		var color c80color.Type = c80color.Type(buffer[:])

		color.Poke(expected[:]...)

		if actual := buffer; expected != actual {
			t.Errorf("For test #%d, actual is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
