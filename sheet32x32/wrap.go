package c80sheet32x32

func Wrap(p []uint8) (Type, error) {
	if nil == p {
		return Nothing(), errNil
	}

	if ByteSize != len(p) {
		return Nothing(), errBadLength
	}

	return Type{p}, nil
}
