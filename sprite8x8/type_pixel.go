package c80sprite8x8

import (
	"github.com/reiver/go-c80/pixel"
)

func (receiver Type) Pixel(x int, y int) c80pixel.Type {
	p := receiver.bytes

	if nil == p {
		return c80pixel.Nothing()
	}
	if 0 >= len(p) {
		return c80pixel.Nothing()
	}
	if ByteSize != len(p) {
		return c80pixel.Nothing()
	}

	offset := offset(x,y)

	if len(p) <= offset {
		return c80pixel.Nothing()
	}

	pixelMemory := p[offset:offset+c80pixel.ByteSize]

	pixel, err := c80pixel.Wrap(pixelMemory)
	if nil != err {
		return c80pixel.Nothing()
	}

	return pixel
}
