package c80sprite

import (
	"github.com/reiver/go-c80/pixel"
)

func (receiver Type) Pixel(x int, y int) c80pixel.Type {
	if nil == receiver {
		return nil
	}

	offset := Offset(x,y)

	if len(receiver) <= offset {
		return nil
	}

	return c80pixel.Type(receiver[offset:offset+c80pixel.Len])
}
