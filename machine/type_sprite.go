package c80machine

import (
	"image"
	"image/draw"
)

// Sprite returns a sprite as a Go built-in image.Image.
//
// There are 3 kinds of sprites:
//
// • 8x8
// • 32x32
// • tile
//
// For each fo these, the id runs from 0 to 255.
func (receiver *Type) Sprite(kind string, id uint8) image.Image {
	if nil == receiver {
		return nil
	}

	switch kind {
	case "8x8":
		return receiver.Sprites8x8().Sprite(id)
	case "32x32":
		return receiver.Sprites32x32().Sprite(id)
	case "tile":
		return receiver.Tiles().Sprite(id)
	default:
		return nil
	}
}

// SetSprite sets the value of a sprite to the built-in image.Image.
//
// There are 3 kinds of sprites:
//
// • 8x8
// • 32x32
// • tile
//
// For each fo these, the id runs from 0 to 255.
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
	case "tile":
		dst = receiver.Tiles().Sprite(id)
	default:
		return errSpriteKindNotAllowed
	}

	{
		rect := img.Bounds()

		draw.Draw(dst, rect, img, rect.Min, draw.Src)
	}

	return nil
}
