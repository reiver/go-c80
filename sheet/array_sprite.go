package c80sheet

import (
	"github.com/reiver/go-c80/sprite"
)

func (receiver *Array) Sprite(n uint8) c80sprite.Type {
	if nil == receiver {
		return nil
	}

	sheet := Type(receiver[:])

	return sheet.Sprite(n)
}
