package c80

import (
	"errors"
)

var (
	errBadRequest    = errors.New("bad request")
	errInternalError = errors.New("internal error")
)
