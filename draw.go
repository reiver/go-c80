package c80

import (
	"github.com/reiver/go-c80/dye"

	"github.com/reiver/go-pel"

	"image"
)

// Draw draws the ‘img’.
//
// Examples
//
// Here are some examples of Draw being used:
//
//	err := c80.Draw(c80.Dye(index))
//
// And:
//
//	err := c80.Draw(c80.Pixel(index, x, y))
//
// And:
//
//	err := c80.Draw(c80.Map(mx, my, width,height))
//
// And:
//
//	err := c80.Draw(c80.Relocate(x,y, c80.Sprite("32x32", id)))
//
// And:
//
//	err := c80.Draw(c80.Text("Hello world!"), x, y)
func Draw(img image.Image) error {
	switch casted := img.(type) {
	case c80dye.Type:
		return machine.DrawDye(casted)
	case *c80dye.Type:
		return machine.DrawDye(casted)

	case pel.RGBA:
		return machine.DrawPixel(casted.X, casted.Y, casted)
	case *pel.RGBA:
		return machine.DrawPixel(casted.X, casted.Y, casted)

	default:
		return machine.Draw(img)
	}
}
