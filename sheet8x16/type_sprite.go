package c80sheet8x16

import (
	"github.com/reiver/go-c80/sprite8x16"
)

func (receiver Type) Sprite(n uint8) c80sprite8x16.Type {
	if nil == receiver {
		return nil
	}

	beginning := offset(n)
	ending    := beginning +c80sprite8x16.Len

	return c80sprite8x16.Type(receiver[beginning:ending])
}
