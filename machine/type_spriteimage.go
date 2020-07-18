package c80machine

import (
	"image"
)

func (receiver *Type) SpriteImage(kind string, id uint8) image.Image {
	if nil == receiver {
		return nil
	}

	switch kind {
	case "8x8":
		return &(receiver.images.sprites8x8[id])
	case "32x32":
		return &(receiver.images.sprites32x32[id])
	default:
		return nil
	}
}
