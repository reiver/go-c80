package c80palette_test

import (
	"github.com/reiver/go-c80/palette"

	"fmt"
)

func ExampleType_Color() {

	// This memory will be used for the palette.
	// This palette will store 256 colors.
	var memory [c80palette.ByteSize]uint8

	var palette c80palette.Type
	var err error

	palette, err = c80palette.Wrap(memory[:])
	if nil != err {
		fmt.Printf("ERROR: could not wrap memory: %s", err)
		return
	}

	// We set the RGBA values for some of the colors in the palette.
	//
	// (“RGBA” = Red, Green, Blue, Alpha.)
	//
	// So, here, color №5 is rgba(118,38,113,255).
	//
	// And, also here, color №0 is rgba(1,1,1,255).
	palette.Color(0).Poke(   1,  1,  1, 255)
	palette.Color(1).Poke( 222, 56, 43, 255)
	palette.Color(2).Poke(  57,181, 74, 255)
	palette.Color(3).Poke( 255,199,  6, 255)
	palette.Color(4).Poke(   0,111,184, 255)
	palette.Color(5).Poke( 118, 38,113, 255)
	palette.Color(6).Poke(  44,181,233, 255)
	palette.Color(7).Poke( 204,204,204, 255)
	palette.Color(8).Poke( 128,128,128, 255)
	palette.Color(9).Poke( 255,  0,  0, 255)
	palette.Color(10).Poke(  0,255,  0, 255)
	palette.Color(11).Poke(255,255,  0, 255)
	palette.Color(12).Poke(  0,  0,255, 255)
	palette.Color(13).Poke(255,  0,255, 255)
	palette.Color(14).Poke(  0,255,255, 255)
	palette.Color(15).Poke(255,255,255, 255)

	fmt.Println(palette.Color(0))
	fmt.Println(palette.Color(1))
	fmt.Println(palette.Color(2))
	fmt.Println(palette.Color(3))
	fmt.Println(palette.Color(4))
	fmt.Println(palette.Color(5))
	fmt.Println(palette.Color(6))
	fmt.Println(palette.Color(7))
	fmt.Println(palette.Color(8))
	fmt.Println(palette.Color(9))
	fmt.Println(palette.Color(10))
	fmt.Println(palette.Color(11))
	fmt.Println(palette.Color(12))
	fmt.Println(palette.Color(13))
	fmt.Println(palette.Color(14))
	fmt.Println(palette.Color(15))

	// Output:
	// rgba(1,1,1,255)
	// rgba(222,56,43,255)
	// rgba(57,181,74,255)
	// rgba(255,199,6,255)
	// rgba(0,111,184,255)
	// rgba(118,38,113,255)
	// rgba(44,181,233,255)
	// rgba(204,204,204,255)
	// rgba(128,128,128,255)
	// rgba(255,0,0,255)
	// rgba(0,255,0,255)
	// rgba(255,255,0,255)
	// rgba(0,0,255,255)
	// rgba(255,0,255,255)
	// rgba(0,255,255,255)
	// rgba(255,255,255,255)
}
