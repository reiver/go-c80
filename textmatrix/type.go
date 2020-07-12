package c80textmatrix

import (
	"reflect"
	"unsafe"
)

// Type represents a text matrix.
type Type []uint8

func (receiver Type) Runes() []rune {
	if nil == receiver {
		return nil
	}

	h := (*reflect.SliceHeader)(unsafe.Pointer(&receiver))

	var header reflect.SliceHeader
	header.Data = h.Data
	header.Len  = h.Len/4
	header.Cap  = h.Cap/4

	return *(*[]rune)(unsafe.Pointer(&header))
}
