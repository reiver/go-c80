package c80frame

import (
	"image"
	"image/color"
	"image/draw"
)

func (receiver *Type) Dye(c color.Color) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == c {
		return errNilColor
	}

	dst := receiver.DrawableImage()
	if nil == dst {
		return errInternalError
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
				X:x+Width,
				Y:y+Height,
			},
		}
	}

	src := &image.Uniform{c}

	draw.Draw(dst, rect, src, image.ZP, draw.Src)

	return nil
}
