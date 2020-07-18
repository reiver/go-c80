package c80frame

import (
	"errors"
)

var (
	errInternalError = errors.New("internal error")
	errNilColor      = errors.New("nil color")
	errNilReceiver   = errors.New("nil receiver")
	errNilWriter     = errors.New("nil writer")
)
