package c80sprite8x16

import (
	"github.com/reiver/go-c80/pixel"
)

// Len represents the number of bytes in a sprite.
//
// Each sprite is 8×16 pixels; which means there are 128 pixels total.
// Each pixel is 1 byte. Therefore the number of bytes in a sprite is 128.
const Len = Width*Height * c80pixel.Len
