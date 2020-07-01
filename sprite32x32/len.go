package c80sprite32x32

import (
	"github.com/reiver/go-c80/pixel"
)

// Len represents the number of bytes in a sprite.
//
// Each sprite is 32×32 pixels; which means there are 1024 pixels total.
// Each pixel is 1 byte. Therefore the number of bytes in a sprite is 1024.
const Len = Width*Height * c80pixel.Len
