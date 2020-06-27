package c80palette

import (
	"github.com/reiver/go-c80/color"
)

// Type represents a (color) palette of 16 colors.
type Type struct {
	colors [Size]c80color.Type
}
