package c80sheet32x32

import (
	"github.com/reiver/go-c80/sprite32x32"
)

func (receiver *Array) Sprite(n uint8) c80sprite32x32.Type {
	if nil == receiver {
		return nil
	}

	sheet := Type(receiver[:])

	return sheet.Sprite(n)
}
