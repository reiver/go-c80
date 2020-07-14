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

		var buffer [4]uint8
		for i:=0; i<len(buffer); i++ {
			buffer[i] = uint8(randomness.Intn(256))
		}

		var expected [4]uint8
		for i:=0; i<len(expected); i++ {
			expected[i] = uint8(randomness.Intn(256))
		}

		color, err := c80color.Wrap(buffer[:])
		if nil != err {
			t.Errorf("For test #%d, received an error, but did not actually expect one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
		}

		color.Poke(expected[0], expected[1], expected[2], expected[3])

		if actual := buffer; expected != actual {
			t.Errorf("For test #%d, actual is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
