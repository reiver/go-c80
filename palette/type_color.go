package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Color returns the color in the palette at ‘index’.
//
// These color palettes only have 256 colors — from 0 to 255.
func (receiver Type) Color(index uint8) (color c80color.Type) {
	if nil == receiver {
		return nil
	}

	defer func() {
		if r := recover(); nil != r {
			color = nil
		}
	}()

	beginning := int(index) * c80color.Len
	ending := beginning + c80color.Len

	{
		length := len(receiver)

		if length <= int(beginning) {
			return nil
		}

		if length < int(ending) {
			return nil
		}
	}

	return c80color.Type(receiver[beginning:ending])
}
