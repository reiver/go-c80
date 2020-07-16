package c80palettedraster_test

import (
	"github.com/reiver/go-c80/palettedraster"

	"math/rand"
	"time"

	"testing"
)

func TestType_IsSomething(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	var buffer [c80palettedraster.ByteSize]uint8
	for i:=0; i<len(buffer); i++ {
		buffer[i] = uint8(randomness.Intn(256))
	}

	var palettedraster c80palettedraster.Type
	var err error

	palettedraster, err = c80palettedraster.Wrap(buffer[:])
	if nil != err {
		t.Errorf("Received an error, but did not actually expect one.")
		t.Logf("ERROR: (%T) %q", err, err)
		return
	}

	{
		expected := true
		actual   := palettedraster.IsSomething()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}

	{
		expected := false
		actual   := palettedraster.IsNothing()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
	}
}
