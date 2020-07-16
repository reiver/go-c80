package c80machine

import (
	"io"
)

func (receiver *Type) PNGTo(writer io.Writer) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == writer {
		return errNilWriter
	}

	if err := receiver.frame.PNGTo(writer); nil != err {
		return err
	}

	return nil
}
