package c80sheet

import (
	"github.com/reiver/go-c80/sprite8x8"
)

func (receiver *Array) Sprite(n uint8) c80sprite8x8.Type {
	if nil == receiver {
		return nil
	}

	sheet := Type(receiver[:])

	return sheet.Sprite(n)
}
