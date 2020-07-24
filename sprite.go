package c80

import (
	"image"
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
func Sprite(kind string, id uint8) image.Image {
	return machine.Sprite(kind, id)
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
func SetSprite(kind string, id uint8, img image.Image) error {
	return machine.SetSprite(kind, id, img)
}
