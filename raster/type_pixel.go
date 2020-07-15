package c80raster

import (
	"github.com/reiver/go-c80/pixel"
)

// Pixel returns the pixel in the raster at (‘x’, ‘y’).
func (receiver Type) Pixel(x int, y int) (pixel c80pixel.Type) {
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
		return c80pixel.Nothing()
	}

	defer func() {
		if r := recover(); nil != r {
			pixel = c80pixel.Nothing()
		}
	}()

	index := receiver.PixOffset(x,y)

	beginning := int(index) * c80pixel.ByteSize
	ending := beginning     + c80pixel.ByteSize

	{
		length := len(p)

		if length <= int(beginning) {
			return c80pixel.Nothing()
		}

		if length < int(ending) {
			return c80pixel.Nothing()
		}
	}

	pixelSlice := p[beginning:ending]

	{
		var err error

		pixel, err = c80pixel.Wrap(pixelSlice)
		if nil != err {
			return c80pixel.Nothing()
		}
	}

	return pixel
}
