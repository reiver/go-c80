package c80sheet8x16

import (
	"github.com/reiver/go-c80/sprite8x16"
)

func offset(n uint8) int {
	return int(n) * c80sprite8x16.Len
}
