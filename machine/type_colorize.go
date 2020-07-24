package c80machine

import (
	"github.com/reiver/go-palette2048"
	"github.com/reiver/go-palette2048_gb"
	"github.com/reiver/go-palette2048_greer"
	"github.com/reiver/go-palette2048_gruvbox"
	"github.com/reiver/go-palette2048_html3"
	"github.com/reiver/go-palette2048_nes"
	"github.com/reiver/go-palette2048_solarized"
	"github.com/reiver/go-palette2048_tia"
)

// Colorize sets the color palette.
//
// Example
//
// Here are some examples of Colorize being used:
//
//	err := machine.Colorize("tia")
func (receiver *Type) Colorize(a ...interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

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
		case "gb":
			p = palette2048_gb.Palette[:]
		case "greer":
			p = palette2048_greer.Palette[:]
		case "gruvbox":
			p = palette2048_gruvbox.Palette[:]
		case "html3":
			p = palette2048_html3.Palette[:]
		case "nes":
			p = palette2048_nes.Palette[:]
		case "solarized":
			p = palette2048_solarized.Palette[:]
		case "tia":
			p = palette2048_tia.Palette[:]
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

	palette := receiver.Palette()
	if nil == palette {
		return errInternalError
	}

	for i:=0; i<len(palette); i++ {
		palette[i] = p[i]
	}

	return nil
}
