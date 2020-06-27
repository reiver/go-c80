package c80

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"
)

var (
	raster c80raster.Type
)

func init() {
	raster.SetPalette(
		c80palette.RGBs(
			c80color.RGB(  1,  1,  1), // black
			c80color.RGB(222, 56, 43), // red
			c80color.RGB( 57,181, 74), // green
			c80color.RGB(255,199,  6), // yellow
			c80color.RGB(  0,111,184), // blue
			c80color.RGB(118, 38,113), // magenta
			c80color.RGB( 44,181,233), // cyan
			c80color.RGB(204,204,204), // white

			c80color.RGB(128,128,128), // bright black
			c80color.RGB(255,  0,  0), // bright red
			c80color.RGB(  0,255,  0), // bright green
			c80color.RGB(255,255,  0), // bright yellow
			c80color.RGB(  0,  0,255), // bright blue
			c80color.RGB(255,  0,255), // bright magenta
			c80color.RGB(  0,255,255), // bright cyan
			c80color.RGB(255,255,255), // bright white
		),
	)
}
