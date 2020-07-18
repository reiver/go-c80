package c80

import (
	"image"
)

func Map(x int, y int, width int, height int) image.Image {
	return machine.MapImage(x, y, width, height)
}
