package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"math/rand"
	"time"

	"testing"
)

func TestWrap(t *testing.T) {

	{
		var p []uint8 = nil

		_, err := c80color.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}



	{
		var buffer [1]uint8

		var p []uint8 = buffer[:]

		_, err := c80color.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [2]uint8

		var p []uint8 = buffer[:]

		_, err := c80color.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [3]uint8

		var p []uint8 = buffer[:]

		_, err := c80color.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}



	{
		var buffer [5]uint8

		var p []uint8 = buffer[:]

		_, err := c80color.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [6]uint8

		var p []uint8 = buffer[:]

		_, err := c80color.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}
	{
		var buffer [7]uint8

		var p []uint8 = buffer[:]

		_, err := c80color.Wrap(p)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one.")
		}
	}



	{
		var buffer [4]uint8

		var p []uint8 = buffer[:]

		_, err := c80color.Wrap(p)
		if nil != err {
			t.Errorf("Did not expect an error, but did actually gott one.")
			t.Logf("ERROR: (%T) %q", err, err)
		}
	}



	for testNumber:=0; testNumber < 20; testNumber++ {
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		var buffer [4]uint8

		for i:=0; i<len(buffer); i++ {
			buffer[i] = uint8(randomness.Intn(256))
		}

		var p []uint8 = buffer[:]

		color, err := c80color.Wrap(p)
		if nil != err {
			t.Errorf("Did not expect an error, but did actually gott one.")
			t.Logf("ERROR: (%T) %q", err, err)
		}

		{
			r,g,b,a := color.Peek()

			if r != buffer[0] || g != buffer[1] || b != buffer[2] || a != buffer[3] {
				t.Errorf("The actual rgba values are not what was expected.")
				t.Logf("EXPECTED rgba: (%d,%d,%d,%d)", buffer[0], buffer[1], buffer[2], buffer[3])
				t.Logf("ACTUAL   rgba: (%d,%d,%d,%d)", r,g,b,a)
			}
		}

		{
			r := uint8(randomness.Intn(256))
			g := uint8(randomness.Intn(256))
			b := uint8(randomness.Intn(256))
			a := uint8(randomness.Intn(256))

			color.Poke(r,g,b,a)

			if r != buffer[0] || g != buffer[1] || b != buffer[2] || a != buffer[3] {
				t.Errorf("The actual rgba values are not what was expected.")
				t.Logf("EXPECTED rgba: (%d,%d,%d,%d)", buffer[0], buffer[1], buffer[2], buffer[3])
				t.Logf("ACTUAL   rgba: (%d,%d,%d,%d)", r,g,b,a)
			}
		}
	}
}
