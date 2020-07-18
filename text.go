package c80

import (
	"github.com/reiver/go-c80/text"

	"image"
)

func Text(msg string) image.Image {
	return c80text.Type{msg}
}
