package c80sheet8x8

import (
	"github.com/reiver/go-c80/sprite8x8"
)

func offset(n uint8) int {
	return int(n) * c80sprite8x8.ByteSize
}
