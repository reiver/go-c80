package c80sheet8x16

import (
	"github.com/reiver/go-c80/sprite8x16"
)

func (receiver *Array) Sprite(n uint8) c80sprite8x16.Type {
	if nil == receiver {
		return nil
	}

	sheet := Type(receiver[:])

	return sheet.Sprite(n)
}
