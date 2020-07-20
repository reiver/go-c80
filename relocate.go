package c80

import (
	"github.com/reiver/go-imagerelocate"

	"image"
)

func Relocate(x, y int, img image.Image) image.Image {
	return imagerelocate.Wrap(x,y, img)
}
