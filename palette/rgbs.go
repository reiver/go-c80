package c80palette

import (
	"github.com/reiver/go-c80/color"
)

func RGBs(rgbs ...c80color.Type) Type {

	var colors [Size]c80color.Type

	copy(colors[:], rgbs)

	return Type{
		colors: colors,
	}
}
