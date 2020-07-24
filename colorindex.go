package c80

import (
	"image/color"
)

func ColorIndex(c color.Color) uint8 {
	return machine.ColorIndex(c)
}
