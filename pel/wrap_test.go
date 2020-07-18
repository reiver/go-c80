package c80pel_test

import (
	"github.com/reiver/go-c80/pel"

	"math/rand"
	"time"

	"testing"
)

func TestWrap(t *testing.T) {

	{
		var p []uint8 = nil

		_, err := c80pel.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}



	{
		var buffer [2]uint8

		var p []uint8 = buffer[:]

		_, err := c80pel.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [3]uint8

		var p []uint8 = buffer[:]

		_, err := c80pel.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [4]uint8

		var p []uint8 = buffer[:]

		_, err := c80pel.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [5]uint8

		var p []uint8 = buffer[:]

		_, err := c80pel.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [6]uint8

		var p []uint8 = buffer[:]

		_, err := c80pel.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [7]uint8

		var p []uint8 = buffer[:]

		_, err := c80pel.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}



	{
		var buffer [1]uint8

		var p []uint8 = buffer[:]

		_, err := c80pel.Wrap(p)
		if nil != err {
			t.Errorf("Did not expect an error, but did actually gott one.")
			t.Logf("ERROR: (%T) %q", err, err)
		}
	}



	for testNumber:=0; testNumber < 20; testNumber++ {
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		var buffer [1]uint8

		for i:=0; i<len(buffer); i++ {
			buffer[i] = uint8(randomness.Intn(256))
		}

		var p []uint8 = buffer[:]

		color, err := c80pel.Wrap(p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but did actually gott one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
		}

		{
			index := color.Peek()

			if index != buffer[0] {
				t.Errorf("For test #%d, the actual index values are not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", buffer[0])
				t.Logf("ACTUAL:   %d", index)
			}
		}

		{
			index := uint8(randomness.Intn(256))

			color.Poke(index)

			if index != buffer[0] {
				t.Errorf("For test #%d, the actual index values are not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", buffer[0])
				t.Logf("ACTUAL:   %d", index)
			}
		}
	}
}
