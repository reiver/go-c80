package c80

import (
	"image"
)

func Sprite(kind string, id uint8) image.Image {
	switch kind {
	case "8x8":
		return machine.Sprites8x8().Sprite(id)
	case "32x32":
		return machine.Sprites32x32().Sprite(id)
	default:
		return nil
	}
}
