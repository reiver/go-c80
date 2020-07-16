package c80frame

import (
	"image"
	"image/draw"
)

func (receiver *Type) DrawableImage() draw.Image {
	if nil == receiver {
		return nil
	}

	return &image.NRGBA{
		Pix: receiver.bytes[:],
		Stride:Width*Depth,
		Rect: image.Rectangle{
			Min: image.Point{
				X:0,
				Y:0,
			},
			Max: image.Point{
				X:Width,
				Y:Height,
			},
		},
	}
}
