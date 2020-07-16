package c80machine

import (
	"errors"
)

var (
	errNilReceiver   = errors.New("nil receiver")
	errNilWriter     = errors.New("nil writer")
)
