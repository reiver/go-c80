package c80textmatrix

import (
	"unsafe"
)

const (
        Width  = 32
        Height = 32
	Depth  = unsafe.Sizeof(rune(0))
)
