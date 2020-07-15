package c80machine

import (
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"
	"github.com/reiver/go-c80/sheet8x8"
	"github.com/reiver/go-c80/sheet32x32"
	"github.com/reiver/go-c80/textmatrix"
)

const (
	PTR_PALETTE      = 0
	LEN_PALETTE      = c80palette.ByteSize

	PTR_TILES        = PTR_PALETTE      + LEN_PALETTE
	LEN_TILES        = c80sheet8x8.ByteSize

	PTR_SPRITES8x8   = PTR_TILES        + LEN_TILES
	LEN_SPRITES8x8   = c80sheet8x8.ByteSize

	PTR_SPRITES32x32 = PTR_SPRITES8x8   + LEN_SPRITES8x8
	LEN_SPRITES32x32 = c80sheet32x32.ByteSize

	PTR_TEXTMATRIX   = PTR_SPRITES32x32 + LEN_SPRITES32x32
	LEN_TEXTMATRIX   = c80textmatrix.ByteSize

	PTR_RASTER0      = PTR_TEXTMATRIX   + LEN_TEXTMATRIX
	LEN_RASTER0      = c80raster.ByteSize
)

const Len =
	LEN_PALETTE      +
	LEN_TILES        +
	LEN_SPRITES8x8   +
	LEN_SPRITES32x32 +
	LEN_TEXTMATRIX   +
	LEN_RASTER0
