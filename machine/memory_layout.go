package c80machine

import (
	"github.com/reiver/go-c80/palette"
	"github.com/reiver/go-c80/raster"
	"github.com/reiver/go-c80/sheet8x8"
	"github.com/reiver/go-c80/sheet8x16"
	"github.com/reiver/go-c80/sheet32x32"
)

const (
	PTR_PALETTE      = 0
	LEN_PALETTE      = c80palette.Len

	PTR_FONTS        = PTR_PALETTE      + LEN_PALETTE
	LEN_FONTS        = c80sheet8x8.Len

	PTR_TILES        = PTR_FONTS        + LEN_FONTS
	LEN_TILES        = c80sheet8x8.Len

	PTR_SPRITES8x8   = PTR_TILES        + LEN_TILES
	LEN_SPRITES8x8   = c80sheet8x8.Len

	PTR_SPRITES8x16  = PTR_SPRITES8x8   + LEN_SPRITES8x8
	LEN_SPRITES8x16  = c80sheet8x16.Len

	PTR_SPRITES32x32 = PTR_SPRITES8x16  + LEN_SPRITES8x16
	LEN_SPRITES32x32 = c80sheet32x32.Len

	PTR_RASTER0      = PTR_SPRITES32x32 + LEN_SPRITES32x32
	LEN_RASTER0      = c80raster.Len

	PTR_RASTER1      = PTR_RASTER0      + LEN_RASTER0
	LEN_RASTER1      = c80raster.Len
)

const Len =
	LEN_PALETTE      +
	LEN_FONTS        +
	LEN_TILES        +
	LEN_SPRITES8x8   +
	LEN_SPRITES8x16  +
	LEN_SPRITES32x32 +
	LEN_RASTER0      +
	LEN_RASTER1
