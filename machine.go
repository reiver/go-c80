package c80

import (
	"github.com/reiver/go-c80/machine"
)

var (
	machine c80machine.Type
)

func init() {
	machine.Init()

	palette := machine.Palette()

	palette.Color( 0).Poke(  1,  1,  1, 255) // black
	palette.Color( 1).Poke(222, 56, 43, 255) // red
	palette.Color( 2).Poke( 57,181, 74, 255) // green
	palette.Color( 3).Poke(255,199,  6, 255) // yellow
	palette.Color( 4).Poke(  0,111,184, 255) // blue
	palette.Color( 5).Poke(118, 38,113, 255) // magenta
	palette.Color( 6).Poke( 44,181,233, 255) // cyan
	palette.Color( 7).Poke(204,204,204, 255) // white

	palette.Color( 8).Poke(128,128,128, 255) // bright black
	palette.Color( 9).Poke(255,  0,  0, 255) // bright red
	palette.Color(10).Poke(  0,255,  0, 255) // bright green
	palette.Color(11).Poke(255,255,  0, 255) // bright yellow
	palette.Color(12).Poke(  0,  0,255, 255) // bright blue
	palette.Color(13).Poke(255,  0,255, 255) // bright magenta
	palette.Color(14).Poke(  0,255,255, 255) // bright cyan
	palette.Color(15).Poke(255,255,255, 255) // bright white
}
