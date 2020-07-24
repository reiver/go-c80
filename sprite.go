package c80

import (
	"image"
)

func Sprite(kind string, id uint8) image.Image {
	return machine.Sprite(kind, id)
}

func SetSprite(kind string, id uint8, img image.Image) error {
	return machine.SetSprite(kind, id, img)
}
