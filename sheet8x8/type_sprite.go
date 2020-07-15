package c80sheet8x8

import (
	"github.com/reiver/go-c80/sprite8x8"
)

func (receiver Type) Sprite(index uint8) (sprite8x8 c80sprite8x8.Type) {
	p := receiver.bytes

	if nil == p {
		return c80sprite8x8.Nothing()
	}

	defer func() {
		if r := recover(); nil != r {
			sprite8x8 = c80sprite8x8.Nothing()
		}
	}()

	beginning := int(index) * c80sprite8x8.ByteSize
	ending := beginning     + c80sprite8x8.ByteSize

	{
		length := len(p)

		if length <= int(beginning) {
			return c80sprite8x8.Nothing()
		}

		if length < int(ending) {
			return c80sprite8x8.Nothing()
		}
	}

	sprite8x8Slice := p[beginning:ending]

	{
		var err error

		sprite8x8, err = c80sprite8x8.Wrap(sprite8x8Slice)
		if nil != err {
			return c80sprite8x8.Nothing()
		}
	}

	return sprite8x8
}
