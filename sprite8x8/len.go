package c80sprite

import (
	"github.com/reiver/go-c80/pixel"
)

// Len represents the number of bytes in a sprite.
//
// Each sprite is 8×8 pixels; which means there are 64 pixels total.
// Each pixel is 1 byte. Therefore the number of bytes in a sprite is 64.
const Len = Width*Height * c80pixel.Len
