package c80sheet32x32

import (
	"github.com/reiver/go-c80/sprite32x32"
)

func (receiver Type) Sprite(index uint8) (sprite32x32 c80sprite32x32.Type) {
	p := receiver.bytes

	if nil == p {
		return c80sprite32x32.Nothing()
	}

	defer func() {
		if r := recover(); nil != r {
			sprite32x32 = c80sprite32x32.Nothing()
		}
	}()

	beginning := int(index) * c80sprite32x32.ByteSize
	ending := beginning     + c80sprite32x32.ByteSize

	{
		length := len(p)

		if length <= int(beginning) {
			return c80sprite32x32.Nothing()
		}

		if length < int(ending) {
			return c80sprite32x32.Nothing()
		}
	}

	sprite32x32Slice := p[beginning:ending]

	{
		var err error

		sprite32x32, err = c80sprite32x32.Wrap(sprite32x32Slice)
		if nil != err {
			return c80sprite32x32.Nothing()
		}
	}

	return sprite32x32
}
