package c80color_test

import (
	"github.com/reiver/go-c80/color"

	"fmt"
)

func ExampleType_String() {

	// This memory will be used to store 16 colors.
	var memory [16*c80color.ByteSize]uint8

	// The colors stored in ‘memory’ are indexed from 0 to 15.
	//
	// Here we are getting the 5th color location in ‘memory’.
	var colorStorage []uint8
	{
		beginning := 5*c80color.ByteSize
		ending    := beginning + c80color.ByteSize

		colorStorage = memory[beginning:ending]
	}

	color, err := c80color.Wrap(colorStorage[:])
	if nil != err {
		fmt.Printf("could not wrap color: %s", err)
		return
	}

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
