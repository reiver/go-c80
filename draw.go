package c80

import (
	"github.com/reiver/go-c80/dye"
	"github.com/reiver/go-c80/pixel"

	"image"
)

// Draw draws the ‘img’ at (‘x’,‘y’).
//
// Examples
//
// Here are some examples of Draw being used:
//
//	err := c80.Draw(c80.Dye(index), 0, 0)
//
// And:
//
//	err := c80.Draw(c80.Pixel(index), x, y)
//
// And:
//
//	err := c80.Draw(c80.Map(mx, my, width,height), x, y)
//
// And:
//
//	err := c80.Draw(c80.Sprite("32x32", id), x, y)
//
// And:
//
//	err := c80.Draw(c80.Text("Hello world!"), x, y)
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
