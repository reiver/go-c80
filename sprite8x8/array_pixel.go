package c80sprite8x8

import (
	"github.com/reiver/go-c80/pixel"
)

func (receiver *Array) Pixel(x int, y int) c80pixel.Type {
	if nil == receiver {
		return nil
	}

	pixel := Type(receiver[:])

	return pixel.Pixel(x,y)
}
