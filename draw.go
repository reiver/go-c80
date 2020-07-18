package c80

import (
	"github.com/reiver/go-c80/dye"
	"github.com/reiver/go-c80/pixel"

	"image"
)

func Draw(img image.Image, x int, y int) error {
	switch casted := img.(type) {
	case c80dye.Type:
		return machine.DrawDye(casted)
	case *c80dye.Type:
		return machine.DrawDye(casted)

	case c80pixel.Type:
		return machine.DrawPixel(x,y, casted)
	case *c80pixel.Type:
		return machine.DrawPixel(x,y, casted)

	default:
		return machine.Draw(img, x, y)
	}
}
