package c80machine

import (
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"
	"github.com/reiver/go-c80/sheet8x8"
	"github.com/reiver/go-c80/sheet8x16"
	"github.com/reiver/go-c80/sheet32x32"
)

// Type represents a fantasy virtual machine.
type Type struct {
	memory [Len]uint8
}

// Palette provides access to the machines palette.
func (receiver *Type) Palette() c80palette.Type {
	beginning := PTR_PALETTE
	ending    := beginning + LEN_PALETTE

	return receiver.memory[beginning:ending]
}

// Fonts provides access to the machine's sprite sheet for 8x8 pixel sprite fonts.
func (receiver *Type) Fonts() c80sheet8x8.Type {
	beginning := PTR_FONTS
	ending    := beginning + LEN_FONTS

	return receiver.memory[beginning:ending]
}

// Sprites8x8 provides access to the machine's sprite sheet for 8x8 pixel sprites.
func (receiver *Type) Sprites8x8() c80sheet8x8.Type {
	beginning := PTR_SPRITES8x8
	ending    := beginning + LEN_SPRITES8x8

	return receiver.memory[beginning:ending]
}

// Sprites8x16 provides access to the machine's sprite sheet for 8x16 pixel sprites.
func (receiver *Type) Sprites8x16() c80sheet8x16.Type {
	beginning := PTR_SPRITES8x16
	ending    := beginning + LEN_SPRITES8x16

	return receiver.memory[beginning:ending]
}

// Sprites32x32 provides access to the machine's sprite sheet for 32x32 pixel sprites.
func (receiver *Type) Sprites32x32() c80sheet32x32.Type {
	beginning := PTR_SPRITES32x32
	ending    := beginning + LEN_SPRITES32x32

	return receiver.memory[beginning:ending]
}

// Raster0 provides access to one of the machine's default raster — raster №0.
func (receiver *Type) Raster0() c80raster.Type {
	beginning := PTR_RASTER0
	ending    := beginning + LEN_RASTER0

	return receiver.memory[beginning:ending]
}

// Raster1 provides access to one of the machine's default raster — raster №1.
func (receiver *Type) Raster1() c80raster.Type {
	beginning := PTR_RASTER1
	ending    := beginning + LEN_RASTER1

	return receiver.memory[beginning:ending]
}

// Raster2 provides access to one of the machine's default raster — raster №2.
func (receiver *Type) Raster2() c80raster.Type {
	beginning := PTR_RASTER2
	ending    := beginning + LEN_RASTER2

	return receiver.memory[beginning:ending]
}
