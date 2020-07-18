package c80

import (
	"image"
)

func Sprite(kind string, id uint8) image.Image {
	return machine.SpriteImage(kind, id)
}
