package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"fmt"
)

func ExampleType_String() {

	// This memory will be used to store 16 colors.
	var memory [16*c80color.Len]uint8

	// The colors stored in ‘memory’ are index from 0 to 15.
	//
	// Here we are getting the 5th color location in ‘memory’.
	var colorStorage []uint8
	{
		beginning := 5*c80color.Len
		ending    := beginning + c80color.Len

		colorStorage = memory[beginning:ending]
	}

	var color c80color.Type = c80color.Type(colorStorage[:])

	const red   = 44
	const green = 181
	const blue  = 233
	const alpha = 255
	color.Poke(red,green,blue,alpha)

	s := color.String()

	fmt.Println(s)

	// Output:
	// rgba(44,181,233,255)
}
