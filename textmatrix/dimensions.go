package c80textmatrix

import (
	"unsafe"
)

const (
        Width  = 32
        Height = 32
	Depth  = int(unsafe.Sizeof(rune(0)))
)
