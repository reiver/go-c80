package c80palette

import (
	"github.com/reiver/go-c80/color"
)

func VGA() Type {
	return RGBs(
		c80color.RGB(  0,  0,  0), // Black
		c80color.RGB(170,  0,  0), // Red
		c80color.RGB(  0,170,  0), // Green
		c80color.RGB(170, 85,  0), // Yellow
		c80color.RGB(  0,  0,170), // Blue
		c80color.RGB(170,  0,170), // Magenta
		c80color.RGB(  0,170,170), // Cyan
		c80color.RGB(170,170,170), // White
		c80color.RGB( 85, 85, 85), // Bright Black
		c80color.RGB(255, 85, 85), // Bright Red
		c80color.RGB( 85,255, 85), // Bright Green
		c80color.RGB(255,255, 85), // Bright Yellow
		c80color.RGB( 85, 85,255), // Bright Blue
		c80color.RGB(255, 85,255), // Bright Magenta
		c80color.RGB( 85,255,255), // Bright Cyan
		c80color.RGB(255,255,255), // Bright White
	)
}
