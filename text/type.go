package c80text

import (
	"github.com/reiver/go-font8x8"

	"image"
	"image/color"
	"unicode/utf8"
)

type Type struct {
	Text string
}

func (receiver Type) At(x, y int) color.Color {

	const fontWidth  = 8
	const fontHeight = 8

	{
		const minY = 0
		const maxY = (fontHeight-1)

		if y < minY || maxY < y {
			return nil
		}
	}

	n := utf8.RuneCountInString(receiver.Text)

	{
		const minX = 0
		maxX := n*fontWidth - 1

		if x < minX || maxX < x {
			return nil
		}
	}

	var img image.Image
	{
		offset := x / fontWidth

		s := receiver.Text

		var r rune
		var size int

		for i:=0; i<(offset+1); i++ {
			r, size = utf8.DecodeRuneInString(s)
			s = s[size:]

			if utf8.RuneError == r {
				return nil
			}
		}

		img = font8x8.Rune(r)
	}

	xx := x - (n-1)*fontWidth

	return img.At(xx, y)
}

func (receiver Type) Bounds() image.Rectangle {

	const x = 0
	const y = 0

	const height = 8 + 1

	var width int
	{
		n := utf8.RuneCountInString(receiver.Text)

		width = n*8 + 1
	}

	return image.Rectangle{
		Min: image.Point{
			X:x,
			Y:y,
		},
		Max: image.Point{
			X:x+width,
			Y:y+height,
		},
	}
}

func (receiver Type) ColorModel() color.Model {
	var doesNotMatterWhichOne font8x8.Type = font8x8.U0041

	return doesNotMatterWhichOne.ColorModel()
}
