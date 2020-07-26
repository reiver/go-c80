package c80machine

import (
	"github.com/reiver/go-tilemap8x8x256x256"

	"image"
)

func (receiver *Type) Map() image.Image {
	if nil == receiver {
		return nil
	}

	return tilemap8x8x256x256.TileMap{
		Map: receiver.tilemap(),
		Tiles: func (id uint8) image.Image {
			return receiver.Tiles().Sprite(id)
		},
	}
}
