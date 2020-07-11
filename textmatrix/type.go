package c80textmatrix

import (
	"unsafe"
)

// Type represents a text matrix.
type Type []uint8

func (receiver Type) Runes() []rune {
	return *(*[]rune)(unsafe.Pointer(&receiver))
}
