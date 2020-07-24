package c80machine

import (
	"image"
	"image/draw"
)

func (receiver *Type) Sprite(kind string, id uint8) image.Image {
	if nil == receiver {
		return nil
	}

	switch kind {
	case "8x8":
		return receiver.Sprites8x8().Sprite(id)
	case "32x32":
		return receiver.Sprites32x32().Sprite(id)
	default:
		return nil
	}
}

func (receiver *Type) SetSprite(kind string, id uint8, img image.Image) error {
	if nil == receiver {
		return errNilReceiver
	}

	var dst draw.Image

	switch kind {
	case "8x8":
		dst = receiver.Sprites8x8().Sprite(id)
	case "32x32":
		dst = receiver.Sprites32x32().Sprite(id)
	default:
		return errSpriteKindNotAllowed
	}

	{
		rect := img.Bounds()

		draw.Draw(dst, rect, img, rect.Min, draw.Src)
	}

	return nil
}
