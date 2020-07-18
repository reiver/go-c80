package c80machine

import (
	"errors"
)

var (
	errInternalError = errors.New("internal error")
	errNilReceiver   = errors.New("nil receiver")
	errNilWriter     = errors.New("nil writer")
)
