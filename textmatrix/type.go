package c80textmatrix

import (
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"
)

// Type represents a text matrix.
type Type struct {
	bytes []uint8
}

func (receiver Type) Clear() {
	runes := receiver.Runes()
	if nil == runes {
		return
	}

	for i:=0; i<len(runes); i++ {
		runes[i] = ' '
	}
}

func (receiver Type) LineFeed() {
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
	rs := receiver.Runes()
	if nil == rs {
		return
	}

	copy(rs, runes)
}

func (receiver Type) Publish(s string) {
	runes := receiver.Runes()
	if nil == runes {
		return
	}

	if "" == s {
		receiver.LineFeed()
		return
	}

	p := s

	Loop: for 0 < len(p) {
		receiver.LineFeed()

		y := Height-1
		for x:=0; x<Width; x++ {
			offset := receiver.runesOffset(x,y)

			r, size := utf8.DecodeRuneInString(p)
			p = p[size:]
			if utf8.RuneError == r  && 0 < size {
				continue
			}
			if utf8.RuneError == r {
				break Loop
			}

			runes[offset] = r
		}

	}
}

func (receiver Type) Runes() []rune {
	p := receiver.bytes
	if nil == p {
		return nil
	}
	if 0 >= len(p) {
		return nil
	}
	if ByteSize != len(p) {
		return nil
	}

	h := (*reflect.SliceHeader)(unsafe.Pointer(&p))

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
