package c80color

import (
	"fmt"
	"strings"
)

// String makes c80color.Array fit the fmt.Stringer interface.
// And thus make it work with the fmt.Print() family of functions.
func (receiver Array) String() string {

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
