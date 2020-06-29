package c80color

import (
	"fmt"
	"strings"
)

// String makes c80color.Type fit the fmt.Stringer interface.
// And thus make it work with the fmt.Print() family of functions.
func (receiver Type) String() string {
	if nil == receiver {
		return "rgba(0x0,0x0,0x0,0x0)"
	}

	var buffer strings.Builder
	{
		buffer.WriteString("rgba(")

		for i,datum := range receiver {
			if 0 != i {
				buffer.WriteRune(',')
			}

			fmt.Fprintf(&buffer, "%d", datum)
		}

		buffer.WriteRune(')')
	}

	return buffer.String()
}