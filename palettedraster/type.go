package c80palettedraster

import (
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"
)

type Type struct {
	bytes []byte
}

func (receiver Type) Palette() c80palette.Type {
	begin := 0
	end   := begin + c80palette.ByteSize

	p := receiver.bytes[begin:end]

	x, err := c80palette.Wrap(p)
	if nil != err {
			return c80palette.Nothing()
	}

	return x
}

func (receiver Type) Raster() c80raster.Type {
	begin := c80palette.ByteSize
	end   := begin + c80raster.ByteSize

	p := receiver.bytes[begin:end]

	x, err := c80raster.Wrap(p)
	if nil != err {
			return c80raster.Nothing()
	}

	return x
}
