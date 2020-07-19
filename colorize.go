package c80

import (
	"github.com/reiver/go-palette2048"
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
		case "vt":
			p = palette_vt[:]
		}
	}
	if nil == p {
		return errInternalError
	}
	if 0 >= len(p) {
		return errInternalError
	}
	if palette2048.ByteSize != len(p) {
		return errInternalError
	}

	palette := machine.Palette()
	if nil == palette {
		return errInternalError
	}

	for i:=0; i<len(palette); i++ {
		palette[i] = p[i]
	}

	return nil
}
