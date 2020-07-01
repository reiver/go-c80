package c80sheet

import (
	"github.com/reiver/go-c80/sprite8x8"
)

func (receiver Type) Sprite(n uint8) c80sprite8x8.Type {
	if nil == receiver {
		return nil
	}

	beginning := offset(n)
	ending    := beginning +c80sprite8x8.Len

	return c80sprite8x8.Type(receiver[beginning:ending])
}
