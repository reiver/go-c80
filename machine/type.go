package c80machine

import (
	"github.com/reiver/go-frame256x288"
	"github.com/reiver/go-palette2048"
	"github.com/reiver/go-spritesheet8x8x256"
	"github.com/reiver/go-spritesheet32x32x256"
	"github.com/reiver/go-text32x36"
)

// Type represents a fantasy virtual machine.
type Type struct {
	memory [memoryByteSize]uint8
}

// Palette provides access to the machines palette.
func (receiver *Type) Palette() palette2048.Slice {
	beginning := PTR_PALETTE
	ending    := beginning + LEN_PALETTE

	p := receiver.memory[beginning:ending]

	return palette2048.Slice(p)
}

// frame provides access the machine's frame.
func (receiver *Type) frame() frame256x288.Slice {
	beginning := PTR_FRAME
	ending    := beginning + LEN_FRAME

	p := receiver.memory[beginning:ending]

	return frame256x288.Slice(p)
}

func (receiver *Type) tilemap() []uint8 {
	beginning := PTR_TILEMAP
	ending    := beginning + LEN_TILEMAP

	p := receiver.memory[beginning:ending]

	return p
}

func (receiver *Type) tiles() []uint8 {
	beginning := PTR_TILES
	ending    := beginning + LEN_TILES

	p := receiver.memory[beginning:ending]

	return p
}

// Tiles provides access to the machine's sprite sheet for 8x8 pixel sprite (background) tiles.
func (receiver *Type) Tiles() spritesheet8x8x256.Paletted {
	p := receiver.tiles()

	return spritesheet8x8x256.Paletted{
		Pix: p,
		Palette: receiver.Palette(),
		Category: "tiles",
	}
}

func (receiver *Type) sprites8x8() []uint8 {
	beginning := PTR_SPRITES8x8
	ending    := beginning + LEN_SPRITES8x8

	p := receiver.memory[beginning:ending]

	return p
}

// Sprites8x8 provides access to the machine's sprite sheet for 8x8 pixel sprites.
func (receiver *Type) Sprites8x8() spritesheet8x8x256.Paletted {
	p := receiver.sprites8x8()

	return spritesheet8x8x256.Paletted{
		Pix: p,
		Palette: receiver.Palette(),
		Category: "sprites8x8",
	}
}

func (receiver *Type) sprites32x32() []uint8 {
	beginning := PTR_SPRITES32x32
	ending    := beginning + LEN_SPRITES32x32

	p := receiver.memory[beginning:ending]

	return p
}

// Sprites32x32 provides access to the machine's sprite sheet for 32x32 pixel sprites.
func (receiver *Type) Sprites32x32() spritesheet32x32x256.Paletted {
	p := receiver.sprites32x32()

	return spritesheet32x32x256.Paletted{
		Pix: p,
		Palette: receiver.Palette(),
		Category: "sprites32x32",
	}
}

// TextMatrix provides access to the machine's text matrix.
func (receiver *Type) TextMatrix() text32x36.Slice {
	beginning := PTR_TEXTMATRIX
	ending    := beginning + LEN_TEXTMATRIX

	p := receiver.memory[beginning:ending]

	return text32x36.Slice(p)
}

