package c80machine_test

import (
	"github.com/reiver/go-c80/machine"

	"image"

	"testing"
)

func TestType_Dye(t *testing.T) {

	var machine c80machine.Type
	machine.Init()

	// Check the initial value for the machine.
	{
		var image image.Image = machine.Image()

		bounds := image.Bounds()
		if (bounds.Max.Y - bounds.Min.Y) < 1 ||  (bounds.Max.X - bounds.Min.X) < 1 {
			t.Errorf("Bad bounds for the image.")
			t.Logf("BOUNDS: %#v", bounds)
			return
		}

		for y:=bounds.Min.Y; y<bounds.Max.Y; y++ {
			for x:=bounds.Min.X; x<bounds.Max.X; x++ {

				color := image.At(x,y)

				aR,aG,aB,aA := color.RGBA()

				const eR = 0
				const eG = 0
				const eB = 0
				const eA = 0

				if eR != aR || eG != aG || eB != aB || eA != aA {
					t.Errorf("For (x,y)=(%d,%d), the actual color is not what was expected.", x,y)
					t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR,eG,eB,eA)
					t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR,aG,aB,aA)
					return
				}
			}
		}

	}

	const r = 129
	const g = 18
	const b = 90
	const a = 255

	const index = 5

	{
		machine.Palette().Color(index).Poke(r,g,b,a)
	}

	{
		color := machine.Palette().Color(index)
		aR, aG, aB, aA := color.RGBA()

		eR := uint32(r) * (0xffff/0xff)
		eG := uint32(g) * (0xffff/0xff)
		eB := uint32(b) * (0xffff/0xff)
		eA := uint32(a) * (0xffff/0xff)

		if eR != aR || eG != aG || eB != aB || eA != aA {
			t.Errorf("The actual color in palette at index=%d is not what was expected.", index)
			t.Logf("EXPECTED (r,g,b,a,)=(%d,%d,%d,%d)", eR, eG, eB, eA)
			t.Logf("ACTUAL   (r,g,b,a,)=(%d,%d,%d,%d)", aR, aG, aB, aA)
			return
		}
	}

	machine.Dye(index)

	// Check the value for the frame after it has been died.
	{
		var image image.Image = machine.Image()

		bounds := image.Bounds()
		if (bounds.Max.Y - bounds.Min.Y) < 1 ||  (bounds.Max.X - bounds.Min.X) < 1 {
			t.Errorf("Bad bounds for the image.")
			t.Logf("BOUNDS: %#v", bounds)
			return
		}

		for y:=bounds.Min.Y; y<bounds.Max.Y; y++ {
			for x:=bounds.Min.X; x<bounds.Max.X; x++ {
				color := image.At(x,y)

				aR,aG,aB,aA := color.RGBA()

				const eR = uint32(r) * (0xffff/0xff)
				const eG = uint32(g) * (0xffff/0xff)
				const eB = uint32(b) * (0xffff/0xff)
				const eA = uint32(a) * (0xffff/0xff)

				if eR != aR || eG != aG || eB != aB || eA != aA {
					t.Errorf("For (x,y)=(%d,%d), the actual color is not what was expected.", x,y)
					t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR,eG,eB,eA)
					t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR,aG,aB,aA)
					return
				}
			}
		}
	}
}
