package c80palettedframe

import (
	"image"
	"image/draw"
)

func (receiver Type) DrawImage() draw.Image {
	if nil == receiver.bytes {
		return nil
	}

	return &image.Paletted{
		Pix: receiver.Raster().Bytes(),
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
		Palette: receiver.Palette().ColorPalette(),
	}
}
