package c80machine

import (
	"io"
)

func (receiver *Type) PNGTo(writer io.Writer) {
	if nil == receiver {
		return
	}

	if nil == writer {
		return
	}

	if err := receiver.frame.PNGTo(writer); nil != err {
		return
	}
}
