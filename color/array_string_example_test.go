package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"fmt"
)

func ExampleArray_String() {

	const red   = 44
	const green = 181
	const blue  = 233
	const alpha = 255

	var color c80color.Array = c80color.Array{red,green,blue,alpha}

	s := color.String()

	fmt.Println(s)

	// Output:
	// rgba(44,181,233,255)
}
