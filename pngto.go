package c80

import (
	"io"
)

func PNGTo(writer io.Writer) error {
	return machine.PNGTo(writer)
}
