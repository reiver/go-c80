package c80machine

import (
	"github.com/reiver/go-font8x8"

	"github.com/reiver/go-c80/palettedraster"
	"github.com/reiver/go-c80/textmatrix"

	"image"
	"image/draw"
)

func (receiver *Type) DrawRaster() {
	if nil == receiver {
		return
	}

	dst := receiver.frame.DrawableImage()
	if nil == dst {
		return
	}

	var rect image.Rectangle
	{
		const x = 0
		const y = 0

		rect = image.Rectangle{
			Min: image.Point{
				X:x,
				Y:y,
			},
			Max: image.Point{
				X:x+c80palettedraster.Width,
				Y:y+c80palettedraster.Height,
			},
		}
	}

	var src image.Image
	{
		palettedraster := receiver.PalettedRaster()
		if palettedraster.IsNothing() {
			return
		}

		src = palettedraster.DrawableImage()
		if nil == src {
			return
		}
	}

	draw.Draw(dst, rect, src, image.ZP, draw.Src)
}

func (receiver *Type) DrawTextMatrix() {
	if nil == receiver {
		return
	}

	textmatrix := receiver.TextMatrix()
	if textmatrix.IsNothing() {
		return
	}

	frameImage := receiver.frame.DrawableImage()
	if nil == frameImage {
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

		draw.Draw(frameImage, rect, src, image.ZP, draw.Src)
	}
}
