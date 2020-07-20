package c80machine

import (
	"github.com/reiver/go-font8x8"

	"github.com/reiver/go-c80/textmatrix"

	"image"
	"image/draw"
)

func (receiver *Type) Draw(img image.Image) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.Frame().Draw(img)
}

func (receiver *Type) DrawTextMatrix() {
	if nil == receiver {
		return
	}

	textmatrix := receiver.TextMatrix()
	if textmatrix.IsNothing() {
		return
	}

	frame := receiver.Frame()
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
		cx := offset % c80textmatrix.Width
		cy := offset / c80textmatrix.Height

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
