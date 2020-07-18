package c80

import (
	"github.com/reiver/go-c80/palette"
)

// Colorize sets the color palette.
//
// Example
//
// Here are some examples of Colorize being used:
//
//	err := c80.Colorize("tia128")
func Colorize(a ...interface{}) error {

	if nil == a {
		return errBadRequest
	}
	if 0 >= len(a) {
		return errBadRequest
	}

	if 1 < len(a) {
		return errBadRequest
	}

	a0 := a[0]

	s0, casted := a0.(string)
	if !casted {
		return errBadRequest
	}

	var p []byte
	{
		switch s0 {
		case "tia128":
			p = palette_tia128[:]
		}
	}
	if nil == p {
		return errInternalError
	}
	if 0 >= len(p) {
		return errInternalError
	}
	if c80palette.ByteSize != len(p) {
		return errInternalError
	}

	palette := machine.Palette()
	if palette.IsNothing() {
		return errInternalError
	}

	palette.Replace(p...)

	return nil
}
