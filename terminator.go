package c80

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
