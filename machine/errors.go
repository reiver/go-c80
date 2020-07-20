package c80machine

import (
	"errors"
)

var (
	errBadRequest    = errors.New("bad request")
	errInternalError = errors.New("internal error")
	errNilReceiver   = errors.New("nil receiver")
	errNilWriter     = errors.New("nil writer")
)
