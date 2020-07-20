package c80machine

import (
	"github.com/reiver/go-font8x8"
	"github.com/reiver/go-text32x36"

	"image"
	"image/draw"
)

func (receiver *Type) DrawTextMatrix() {
	if nil == receiver {
		return
	}

	textmatrix := receiver.TextMatrix()
	if nil == textmatrix {
		return
	}

	frame := receiver.frame()
	if nil == frame {
		return
	}

	runes := textmatrix.Runes()
	if nil == runes {
		return
	}

	const fontWidth  = 8
	const fontHeight = 8

	for offset, character := range runes {
		cx := offset % text32x36.Width
		cy := offset / text32x36.Height

		x := cx*fontWidth
		y := cy*fontHeight

		rect := image.Rectangle{
			Min: image.Point{
				X:x,
				Y:y,
			},
			Max: image.Point{
				X:x+fontWidth,
				Y:y+fontHeight,
			},
		}

		var src image.Image = font8x8.Rune(character)

		draw.Draw(frame, rect, src, image.ZP, draw.Src)
	}
}
