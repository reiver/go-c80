package c80machine

import (
	"github.com/reiver/go-frame256x288"
	"github.com/reiver/go-palette2048"
	"github.com/reiver/go-spritesheet8x8x256"
	"github.com/reiver/go-spritesheet32x32x256"
	"github.com/reiver/go-text32x36"
)

const (
	PTR_PALETTE      = 0
	LEN_PALETTE      = palette2048.ByteSize

	PTR_FRAME        = PTR_PALETTE      + LEN_PALETTE
	LEN_FRAME        = frame256x288.ByteSize

	PTR_TILES        = PTR_FRAME        + LEN_FRAME
	LEN_TILES        = spritesheet8x8x256.ByteSize

	PTR_SPRITES8x8   = PTR_TILES        + LEN_TILES
	LEN_SPRITES8x8   = spritesheet8x8x256.ByteSize

	PTR_SPRITES32x32 = PTR_SPRITES8x8   + LEN_SPRITES8x8
	LEN_SPRITES32x32 = spritesheet32x32x256.ByteSize

	PTR_TEXTMATRIX   = PTR_SPRITES32x32 + LEN_SPRITES32x32
	LEN_TEXTMATRIX   = text32x36.ByteSize
)

const memoryByteSize =
	LEN_PALETTE      +
	LEN_FRAME        +
	LEN_TILES        +
	LEN_SPRITES8x8   +
	LEN_SPRITES32x32 +
	LEN_TEXTMATRIX
