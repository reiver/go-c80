package c80machine

import (
	"github.com/reiver/go-c80/frame"
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/palettedraster"
	"github.com/reiver/go-c80/raster"
	"github.com/reiver/go-c80/sheet8x8"
	"github.com/reiver/go-c80/sheet32x32"
	"github.com/reiver/go-c80/textmatrix"

	"image"
)

// Type represents a fantasy virtual machine.
type Type struct {
	frame c80frame.Type
	memory [memoryByteSize]uint8

	tiles        [c80sheet8x8.Len]image.Paletted
	sprites8x8   [c80sheet8x8.Len]image.Paletted
	sprites32x32 [c80sheet32x32.Len]image.Paletted
}

// Init initialized Type. Type needs to be initialized before it is used.
func (receiver *Type) Init() {
	if nil == receiver {
		return
	}

	{
		tiles := receiver.Tiles()

		for id:=0; id<len(receiver.tiles); id++ {
			sprite := tiles.Sprite(uint8(id))

			pix    := sprite.Pix()
			stride := sprite.Stride()

			receiver.tiles[id].Pix     = pix
			receiver.tiles[id].Stride  = stride
		}
	}

	{
		sprites8x8 := receiver.Sprites8x8()

		for id:=0; id<len(receiver.sprites8x8); id++ {
			sprite := sprites8x8.Sprite(uint8(id))

			pix    := sprite.Pix()
			stride := sprite.Stride()

			receiver.sprites8x8[id].Pix     = pix
			receiver.sprites8x8[id].Stride  = stride
		}
	}

	{
		sprites32x32 := receiver.Sprites32x32()

		for id:=0; id<len(receiver.sprites32x32); id++ {
			sprite := sprites32x32.Sprite(uint8(id))

			pix    := sprite.Pix()
			stride := sprite.Stride()

			receiver.sprites32x32[id].Pix     = pix
			receiver.sprites32x32[id].Stride  = stride
		}
	}
}


func (receiver *Type) PalettedRaster() c80palettedraster.Type {
	if nil == receiver {
		return c80palettedraster.Nothing()
	}

	beginning := PTR_PALETTE
	ending    := beginning + LEN_PALETTE + LEN_RASTER

	p := receiver.memory[beginning:ending]

	x, err := c80palettedraster.Wrap(p)
	if nil != err {
		return c80palettedraster.Nothing()
	}

	return x
}

// Palette provides access to the machines palette.
func (receiver *Type) Palette() c80palette.Type {
	beginning := PTR_PALETTE
	ending    := beginning + LEN_PALETTE

	p := receiver.memory[beginning:ending]

	x, err := c80palette.Wrap(p)
	if nil != err {
		return c80palette.Nothing()
	}

	return x
}

// Tiles provides access to the machine's sprite sheet for 8x8 pixel sprite (background) tiles.
func (receiver *Type) Tiles() c80sheet8x8.Type {
	beginning := PTR_TILES
	ending    := beginning + LEN_TILES

	p := receiver.memory[beginning:ending]

	x, err := c80sheet8x8.Wrap(p)
	if nil != err {
		return c80sheet8x8.Nothing()
	}

	return x
}

// Sprites8x8 provides access to the machine's sprite sheet for 8x8 pixel sprites.
func (receiver *Type) Sprites8x8() c80sheet8x8.Type {
	beginning := PTR_SPRITES8x8
	ending    := beginning + LEN_SPRITES8x8

	p := receiver.memory[beginning:ending]

	x, err := c80sheet8x8.Wrap(p)
	if nil != err {
		return c80sheet8x8.Nothing()
	}

	return x
}

// Sprites32x32 provides access to the machine's sprite sheet for 32x32 pixel sprites.
func (receiver *Type) Sprites32x32() c80sheet32x32.Type {
	beginning := PTR_SPRITES32x32
	ending    := beginning + LEN_SPRITES32x32

	p := receiver.memory[beginning:ending]

	x, err := c80sheet32x32.Wrap(p)
	if nil != err {
		return c80sheet32x32.Nothing()
	}

	return x
}

// TextMatrix provides access to the machine's text matrix.
func (receiver *Type) TextMatrix() c80textmatrix.Type {
	beginning := PTR_TEXTMATRIX
	ending    := beginning + LEN_TEXTMATRIX

	p := receiver.memory[beginning:ending]

	x, err := c80textmatrix.Wrap(p)
	if nil != err {
		return c80textmatrix.Nothing()
	}

	return x
}

// Raster provides access the machine's default raster.
func (receiver *Type) Raster() c80raster.Type {
	beginning := PTR_RASTER
	ending    := beginning + LEN_RASTER

	p := receiver.memory[beginning:ending]

	x, err := c80raster.Wrap(p)
	if nil != err {
		return c80raster.Nothing()
	}

	return x
}
