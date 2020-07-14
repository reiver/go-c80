package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Color returns the color in the palette at ‘index’.
//
// These color palettes only have 256 colors — from 0 to 255.
func (receiver Type) Color(index uint8) (color c80color.Type) {
	p := receiver.bytes

	if nil == p {
		return c80color.Nothing()
	}

	defer func() {
		if r := recover(); nil != r {
			color = c80color.Nothing()
		}
	}()

	beginning := int(index) * c80color.ByteSize
	ending := beginning     + c80color.ByteSize

	{
		length := len(p)

		if length <= int(beginning) {
			return c80color.Nothing()
		}

		if length < int(ending) {
			return c80color.Nothing()
		}
	}

	colorSlice := p[beginning:ending]

	{
		var err error

		color, err = c80color.Wrap(colorSlice)
		if nil != err {
			return c80color.Nothing()
		}
	}

	return color
}
