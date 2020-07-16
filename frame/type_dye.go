package c80frame

import (
	"github.com/reiver/go-c80/color"

	"image"
	"image/draw"
)

func (receiver *Type) Dye(color c80color.Type) {
	if color.IsNothing() {
		return
	}

	dst := receiver.DrawableImage()
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
				X:x+Width,
				Y:y+Height,
			},
		}
	}

	src := &image.Uniform{color}

	draw.Draw(dst, rect, src, image.ZP, draw.Src)
}
