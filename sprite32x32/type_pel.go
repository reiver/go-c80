package c80sprite32x32

import (
	"github.com/reiver/go-c80/pel"
)

func (receiver Type) Pel(x int, y int) c80pel.Type {
	p := receiver.bytes

	if nil == p {
		return c80pel.Nothing()
	}
	if 0 >= len(p) {
		return c80pel.Nothing()
	}
	if ByteSize != len(p) {
		return c80pel.Nothing()
	}

	offset := offset(x,y)

	if len(p) <= offset {
		return c80pel.Nothing()
	}

	pelMemory := p[offset:offset+c80pel.ByteSize]

	pel, err := c80pel.Wrap(pelMemory)
	if nil != err {
		return c80pel.Nothing()
	}

	return pel
}
