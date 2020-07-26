package c80machine

import (
	"image"
)

type Terminator interface {
	image.Image

	Clear()
	Height() int
	LineFeed()
	Publish(s string)
	Runes() []rune
	String() string
	Width() int
}
