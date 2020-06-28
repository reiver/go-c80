package c80machine

import (
	"github.com/reiver/go-c80/palette"
)

func (receiver *Type) SetPalette(palette c80palette.Type) {
	receiver.palette = palette
}
