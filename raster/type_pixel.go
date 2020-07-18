package c80raster

import (
	"github.com/reiver/go-c80/pel"
)

// Pel returns the pel in the raster at (‘x’, ‘y’).
func (receiver Type) Pel(x int, y int) (pel c80pel.Type) {
	x = x % Width
	if 0 > x {
		x += Width
	}

	y = y % Height
	if 0 > y {
		y += Height
	}

	p := receiver.bytes

	if nil == p {
		return c80pel.Nothing()
	}

	defer func() {
		if r := recover(); nil != r {
			pel = c80pel.Nothing()
		}
	}()

	index := receiver.PixOffset(x,y)

	beginning := int(index) * c80pel.ByteSize
	ending := beginning     + c80pel.ByteSize

	{
		length := len(p)

		if length <= int(beginning) {
			return c80pel.Nothing()
		}

		if length < int(ending) {
			return c80pel.Nothing()
		}
	}

	pelSlice := p[beginning:ending]

	{
		var err error

		pel, err = c80pel.Wrap(pelSlice)
		if nil != err {
			return c80pel.Nothing()
		}
	}

	return pel
}
