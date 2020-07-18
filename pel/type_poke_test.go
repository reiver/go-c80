package c80pel_test

import (
	"github.com/reiver/go-c80/pel"

	"math/rand"
	"time"

	"testing"
)

func TestType_Poke(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<64; testNumber++ {

		var buffer [1]uint8 = [1]uint8{
			uint8(randomness.Intn(16)),
		}

		var pel c80pel.Type
		var err error

		pel, err = c80pel.Wrap(buffer[:])
		if nil != err {
			t.Errorf("For test #%d, ", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}

		expected := uint8(randomness.Intn(16))

		pel.Poke(expected)

		actual := pel.Peek()

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
