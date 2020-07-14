package c80color

func Wrap(p []uint8) (Type, error) {
	if nil == p {
		return Type{}, errNil
	}

	if 4 != len(p) {
		return Type{}, errBadLength
	}

	return Type{p}, nil
}