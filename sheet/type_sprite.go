package c80sheet

import (
	"github.com/reiver/go-c80/sprite"
)

func (receiver Type) Sprite(n uint8) c80sprite.Type {
	if nil == receiver {
		return nil
	}

	beginning := offset(n)
	ending    := beginning +c80sprite.Len

	return c80sprite.Type(receiver[beginning:ending])
}
