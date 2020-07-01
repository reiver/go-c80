package c80sheet32x32

import (
	"github.com/reiver/go-c80/sprite32x32"
)

func (receiver Type) Sprite(n uint8) c80sprite32x32.Type {
	if nil == receiver {
		return nil
	}

	beginning := offset(n)
	ending    := beginning +c80sprite32x32.Len

	return c80sprite32x32.Type(receiver[beginning:ending])
}
