package c80frame

import (
	"image"
)

func (receiver *Type) Image() image.Image {
	if nil == receiver {
		return nil
	}

	array := receiver.bytes

	return &image.NRGBA{
		Pix: array[:],
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
