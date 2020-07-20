package c80machine_test

import (
	"github.com/reiver/go-c80/machine"

	"fmt"
)

func ExampleType_Palette() {

	var machine c80machine.Type

	// We set the RGBA values of each of the colors in the palette.
	//
	// (“RGBA” = Red, Green, Blue, Alpha.)
	//
	// The colors in the palette are indexed from 0 to 15.
	//
	// So, here, color №5 is rgba(118,38,113,255).
	//
	// And, also here, color №0 is rgba(1,1,1,255).
	machine.Palette().SetColorRGB(0,     1,  1,  1)
	machine.Palette().SetColorRGB(1,   222, 56, 43)
	machine.Palette().SetColorRGB(2,    57,181, 74)
	machine.Palette().SetColorRGB(3,   255,199,  6)
	machine.Palette().SetColorRGB(4,     0,111,184)
	machine.Palette().SetColorRGB(5,   118, 38,113)
	machine.Palette().SetColorRGB(6,    44,181,233)
	machine.Palette().SetColorRGB(7,   204,204,204)
	machine.Palette().SetColorRGB(8,   128,128,128)
	machine.Palette().SetColorRGB(9,   255,  0,  0)
	machine.Palette().SetColorRGB(10,    0,255,  0)
	machine.Palette().SetColorRGB(11,  255,255,  0)
	machine.Palette().SetColorRGB(12,    0,  0,255)
	machine.Palette().SetColorRGB(13,  255,  0,255)
	machine.Palette().SetColorRGB(14,    0,255,255)
	machine.Palette().SetColorRGB(15,  255,255,255)

	fmt.Println(machine.Palette().Color(0))
	fmt.Println(machine.Palette().Color(1))
	fmt.Println(machine.Palette().Color(2))
	fmt.Println(machine.Palette().Color(3))
	fmt.Println(machine.Palette().Color(4))
	fmt.Println(machine.Palette().Color(5))
	fmt.Println(machine.Palette().Color(6))
	fmt.Println(machine.Palette().Color(7))
	fmt.Println(machine.Palette().Color(8))
	fmt.Println(machine.Palette().Color(9))
	fmt.Println(machine.Palette().Color(10))
	fmt.Println(machine.Palette().Color(11))
	fmt.Println(machine.Palette().Color(12))
	fmt.Println(machine.Palette().Color(13))
	fmt.Println(machine.Palette().Color(14))
	fmt.Println(machine.Palette().Color(15))

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
