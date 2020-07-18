package c80

import (
	"image"
)

func Draw(img image.Image, x int, y int) error {
	return machine.Draw(img, x, y)
}
