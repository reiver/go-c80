package c80

import (
	"image"
)

func Sprite(kind string, id uint8) image.Image {
	return machine.Sprite(kind, id)
}

func DrawSprite(t string, id uint8, x int, y int) {
	machine.DrawSprite(t, id, x, y)
}
