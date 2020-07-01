package c80sheet32x32

import (
	"github.com/reiver/go-c80/sprite32x32"
)

func offset(n uint8) int {
	return int(n) * c80sprite32x32.Len
}
