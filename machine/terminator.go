package c80machine

import (
	"image"
)

type Terminator interface {
	image.Image

	Clear()
	LineFeed()
	Publish(s string)
	Runes() []rune
	String() string
}
