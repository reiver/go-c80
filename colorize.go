package c80

import (
	"github.com/reiver/go-c80/palette"
)

func Colorize(a ...interface{}) {

	if nil == a {
		return
	}
	if 0 >= len(a) {
		return
	}

	if 1 < len(a) {
		return
	}

	a0 := a[0]

	s0, casted := a0.(string)
	if !casted {
		return
	}

	var p []byte
	{
		switch s0 {
		case "tia128":
			p = palette_tia128[:]
		}
	}
	if nil == p {
		return
	}
	if 0 >= len(p) {
		return
	}
	if c80palette.ByteSize != len(p) {
		return
	}

	palette := machine.Palette()
	if palette.IsNothing() {
		return
	}

	palette.Replace(p...)
}
