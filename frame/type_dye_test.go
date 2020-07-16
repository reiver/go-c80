package c80frame_test

import (
	"github.com/reiver/go-c80/color"
	"github.com/reiver/go-c80/frame"

	"image"

	"testing"
)

func TestType_Dye(t *testing.T) {

	var frame c80frame.Type

	// Check the initial value for the frame.
	{
		var image image.Image = frame.DrawableImage()

		for y:=0; y<c80frame.Height; y++ {
			for x:=0; x<c80frame.Width; x++ {

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

	var color c80color.Type
	{
		var buffer [c80color.ByteSize]uint8

		var err error

		color, err = c80color.Wrap(buffer[:])
		if nil != err {
			t.Errorf("This should not happen.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}

		color.Poke(r,g,b,a)
	}

	frame.Dye(color)

	// Check the value for the frame after it has been died.
	{
		var image image.Image = frame.DrawableImage()

		for y:=0; y<c80frame.Height; y++ {
			for x:=0; x<c80frame.Width; x++ {

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
