package c80sheet

import (
	"github.com/reiver/go-c80/sprite"
)

func offset(n uint8) int {
	return int(n) * c80sprite.Len
}
