package c80pixel

import (
	"errors"
)

var (
	errNil       = errors.New("c80: nil")
	errBadLength = errors.New("c80: bad length")
)
