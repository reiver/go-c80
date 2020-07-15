package c80frame

import (
	"image"
	"image/draw"
)

func (receiver *Type) DrawImage() draw.Image {
	if nil == receiver {
		return nil
	}

	p := receiver.bytes[:]

	return &image.NRGBA{
		Pix: p,
		Stride:Width*Depth,
		Rect: image.Rectangle{
			Min: image.Point{
				X:0,
				Y:0,
			},
			Max: image.Point{
				X:(Width-1),
				Y:(Height-1),
			},
		},
	}
}
