package c80frame

import (
	"image"
	"image/png"
	"io"
)

func (receiver *Type) PNGTo(writer io.Writer) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == writer {
		return errNilWriter
	}

	var img image.Image = receiver.DrawableImage()
	if nil == img {
		return errInternalError
	}

	err := png.Encode(writer, img)
	if nil != err {
		return err
	}

	return nil
}
