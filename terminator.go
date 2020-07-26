package c80

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
