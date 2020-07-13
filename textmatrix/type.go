package c80textmatrix

import (
	"reflect"
	"strings"
	"unsafe"
)

// Type represents a text matrix.
type Type []uint8

func (receiver Type) Clear() {
	if nil == receiver {
		return
	}

	runes := receiver.Runes()
	if nil == runes {
		return
	}

	for i:=0; i<len(runes); i++ {
		runes[i] = ' '
	}
}

func (receiver Type) LineFeed() {
	if nil == receiver {
		return
	}

	runes := receiver.Runes()
	if nil == runes {
		return
	}

	for y:=1; y<=Height; y++ {
		for x:=0; x<Width; x++ {
			srcOffset := receiver.runesOffset(x,y)

			dstOffset := receiver.runesOffset(x,y-1)

			runes[dstOffset] = runes[srcOffset]
		}
	}

	{
		y := Height-1

		for x:=0; x<Width; x++ {
			offset := receiver.runesOffset(x,y)

			runes[offset] = ' '
		}
	}
}

func (receiver Type) Poke(runes ...rune) {
	if nil == receiver {
		return
	}

	rs := receiver.Runes()
	if nil == rs {
		return
	}

	copy(rs, runes)
}

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

func (receiver Type) String() string {
	runes := receiver.Runes()

	if nil == runes {
		return  ""+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "
	}

	var buffer strings.Builder

	for offset:=0; offset<len(runes); offset++ {
		if x := offset % Width; 0 == x && 0 != offset {
			buffer.WriteRune('\n')
		}

		buffer.WriteRune(runes[offset])
	}

	return buffer.String()
}

func (receiver Type) runesOffset(x int, y int) int {
	x = x % Width
	if 0 > x {
		x += Width
	}

	y = y % Height
	if 0 > y {
		y += Height
	}

	return y*Width + x
}
