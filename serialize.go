package c80

import (
	"fmt"
	"strings"
)

func Serialize() string {
	var buffer strings.Builder

	if err := PNGTo(&buffer); nil != err {
		return fmt.Sprintf("ERROR: could not write PNG data: %s", err)
	}

	return buffer.String()
}
