package c80palettedraster_test

import (
	"github.com/reiver/go-font8x8"

	"github.com/reiver/go-c80/palettedraster"

	"image"
	"image/color"
	"image/draw"

	"testing"
)

func TestType_DrawableImage(t *testing.T) {

	var buffer [c80palettedraster.ByteSize]uint8
	var palettedraster c80palettedraster.Type
	{
		var err error

		palettedraster, err = c80palettedraster.Wrap(buffer[:])
		if nil != err {
			t.Errorf("Received an error, but did not actually expect one.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}
	}

	var bytes []byte
	{
		bytes = palettedraster.Bytes()

		if length := len(bytes); 0 >= length {
			t.Errorf("The actual number of bytes is less than expected.")
			t.Logf("LENGTH: %d", length)
			return
		}
		if expected, actual := c80palettedraster.ByteSize, len(bytes); expected != actual {
			t.Errorf("The actual number of bytes is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	// Check to see what is there BEFORE we draw anything is what we expect.
	for offset, value := range bytes {
		if expected, actual := byte(0), value; expected != actual {
			t.Errorf("For offset=%d, the actual value for the pixel color channel is not what was expected.", offset)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	// Set the palette.
	{
		palette := palettedraster.Palette()

		palette.Color(0).Poke(0x00,0x00,0x00, 0xff) // black
		palette.Color(1).Poke(0x00,0x00,0xff, 0xff) // blue
		palette.Color(2).Poke(0x00,0xff,0x00, 0xff) // green
		palette.Color(3).Poke(0x00,0xff,0xff, 0xff) // cyan
		palette.Color(4).Poke(0xff,0x00,0x00, 0xff) // red
		palette.Color(5).Poke(0xff,0x00,0xff, 0xff) // purple
		palette.Color(6).Poke( 255, 199,   6, 0xff) // yellow
		palette.Color(7).Poke(0xff,0xff,0xff, 0xff) // white
	}

	// 1st Background color.
	bg1 := struct{
		Red  uint8
		Green uint8
		Blue  uint8
		Alpha uint8
	}{
		Red:   0,
		Green: 0,
		Blue:  0,
		Alpha: 255,
	}

	// Set every pixel to the 1st background color.
	{
		drawable := palettedraster.DrawableImage()

		c := color.NRGBA{
			R:bg1.Red,
			G:bg1.Green,
			B:bg1.Blue,
			A:bg1.Alpha,
		}

		for y:=0; y<c80palettedraster.Height; y++ {
			for x:=0; x<c80palettedraster.Width; x++ {
				drawable.Set(x,y, c)
			}
		}
	}

	// Check to see what is there AFTER the 1st time we draw anything is what we expect.
	{
		image := palettedraster.DrawableImage()

		for y:=0; y<c80palettedraster.Height; y++ {
			for x:=0; x<c80palettedraster.Width; x++ {
				color := image.At(x,y)

				r,g,b,a := color.RGBA()

				eR := uint32(bg1.Red)   * (0xffff/0xff)
				eG := uint32(bg1.Green) * (0xffff/0xff)
				eB := uint32(bg1.Blue)  * (0xffff/0xff)
				eA := uint32(bg1.Alpha) * (0xffff/0xff)

				if eR != r || eG != g || eB != b || eA != a {
					t.Errorf("For (x,y)=(%d,%d), the actual rgba value for the pixel is not what was expected.", x,y)
					t.Logf("EXPECTED (r,g,b,a) = (%d,%d,%d,%d)", eR,eG,eB,eA)
					t.Logf("ACTUAL   (r,g,b,a) = (%d,%d,%d,%d)", r,g,b,a)
					t.Log(palettedraster.String())
					return
				}
			}
		}
	}

	// 2nd Background color.
	bg2 := struct{
		Red  uint8
		Green uint8
		Blue  uint8
		Alpha uint8
	}{
		Red:   255,
		Green: 199,
		Blue:  6,
		Alpha: 255,
	}

	for offset:=0; offset<(len(bytes)/4); offset++ {
		cx := offset % 32
		cy := offset / 32

		x := cx*8
		y := cy*8

		drawableImage := palettedraster.DrawableImage()

		rect := image.Rectangle{
			Min: image.Point{
				X:x,
				Y:y,
			},
			Max: image.Point{
				X:x+8,
				Y:y+8,
			},
		}

		c := color.RGBA{
			R:bg2.Red,
			G:bg2.Green,
			B:bg2.Blue,
			A:bg2.Alpha,
		}

		draw.Draw(drawableImage, rect, &image.Uniform{c}, image.ZP, draw.Src)
	}

	// Check to see what is there AFTER the 2nd time we draw anything is what we expect.
	{
		drawable := palettedraster.DrawableImage()

		for y:=0; y<c80palettedraster.Height; y++ {
			for x:=0; x<c80palettedraster.Width; x++ {

				aR, aG, aB, aA := drawable.At(x,y).RGBA()

				eR := uint32(bg2.Red)   * (0xffff/0xff)
				eG := uint32(bg2.Green) * (0xffff/0xff)
				eB := uint32(bg2.Blue)  * (0xffff/0xff)
				eA := uint32(bg2.Alpha) * (0xffff/0xff)

				if eR != aR || eG != aG || eB != aB || eA != aA {
					t.Errorf("For (x,y)=(%d,%d), the actual value for the pixel color is not what was expected.", x,y)
					t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, aA)
					t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
					t.Log(palettedraster.String())
					return
				}
			}
		}
	}

	const fontWidth  = 8
	const fontHeight = 8

	const printX = 10
	const printY = 20

	// Draw the letter "N".
	{
		drawable := palettedraster.DrawableImage()

		var character rune = 'N'

		var src image.Image = font8x8.Rune(character)

		rect := image.Rectangle{
			Min: image.Point{
				X:printX,
				Y:printY,
			},
			Max: image.Point{
				X:printX+fontWidth,
				Y:printY+fontHeight,
			},
		}

		draw.Draw(drawable, rect, src, image.ZP, draw.Src)
	}

	// Check to see what is there AFTER the 3nd time we draw anything is what we expect.
	{
		drawable := palettedraster.DrawableImage()

		for y:=0; y<c80palettedraster.Height; y++ {
			for x:=0; x<c80palettedraster.Width; x++ {

				// If we are on the spot where we drew the character.
				if (printY <= y && y < (printY+fontHeight)) && (printX <= x && x < (printX+fontWidth)) {
					aR, aG, aB, aA := drawable.At(x,y).RGBA()

					e1R, e1G, e1B, e1A := uint32(0x0000),uint32(0x0000),uint32(0x0000), uint32(0xffff)
					e2R, e2G, e2B, e2A := uint32(0xffff),uint32(0xffff),uint32(0xffff), uint32(0xffff)

					if (e1R != aR && e2R != aR) || (e1G != aG && e2G != aG) || (e1B != aB && e2B != aB) || (e1A != aA && e2A != aA) {
						t.Errorf("For (x,y)=(%d,%d), the actual value for the pixel color is not what was expected.", x,y)
						t.Logf("EXPECTED 1 (r,g,b,a)=(%d,%d,%d,%d)", e1R, e1G, e1B, e1A)
						t.Logf("EXPECTED 2 (r,g,b,a)=(%d,%d,%d,%d)", e2R, e2G, e2B, e2A)
						t.Logf("ACTUAL     (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
						return
					}
					continue
				}

				aR, aG, aB, aA := drawable.At(x,y).RGBA()

				eR := uint32(bg2.Red)   * (0xffff/0xff)
				eG := uint32(bg2.Green) * (0xffff/0xff)
				eB := uint32(bg2.Blue)  * (0xffff/0xff)
				eA := uint32(bg2.Alpha) * (0xffff/0xff)

				if eR != aR || eG != aG || eB != aB || eA != aA {
					t.Errorf("For (x,y)=(%d,%d), the actual value for the pixel color is not what was expected.", x,y)
					t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, aA)
					t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
					t.Log(palettedraster.String())
					return
				}
			}
		}
	}

	{
		expected := "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAMAAABsAF1iAAADAFBMVEUAAAAAAP8A/wAA////AAD/AP//xwb///8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB9MpZoAAABAHRSTlP//////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJj4ZeAAAAmtJREFUeJzs1DFqgzEQBtGtsve/cSAEjO3+f8XOVCo/PZDm53gB6AG6APQAXQB6gC4APUAXgB6gC0AP0AWgB+gC0AN0AegBugBex92Z2T1G8gaw5wH2OMD8Gcg5z/fxB5wHuP4EbgPMf3LO8x277ncB6AG6APQAXQB6gC4APUAXgB6gC0AP0AWgB+gC0AN0AegBugD0AF0AeoAuAD1AF4AeoAtAD9AFoAfoAtADdAHoAboA9ABdAHqALgA9QBeAHqALQA/QBaAH6ALQA3QB6AG6APQAXQB6gC4APUAXgB6gC0AP0AWgB+gC0AN0AegBugD0AF0AeoAuAD1AF4AeoAtAD9AFoAfoAtADdAHoAboA9ABdAHqALgA9QBeAHqALQA/QBaAH6ALQA3QB6AG6APQAXQB6gC4APUAXgB6gC0AP0AWgB+gC0AN0AegBugD0AF0AeoAuAD1AF4AeoAtAD9AFoAfoAtADdAHoAboA9ABdAHqALgA9QBeAHqALQA/QBaAH6ALQA3QB6AG6APQAXQB6gC4APUAXgB6gC0AP0AWgB+gC0AN0AegBugD0AF0AeoAuAD1AF4AeoAtAD9AFoAfoAtADdAHoAboA9ABdAHqALgA9QBeAHqALQA/QBaAH6ALQA3QB6AG6APQAXQB6gC4APUAXgB6gC0AP0AWgB+gC0AN0AegBugD0AF0AeoAuAD1AF4AeoAtAD9AFoAfoAtADdAHoAboA9ABdAHqALgA9QBeAHqALQA/QBaAH6ALQA3QB6AG6APQAXQB6gC4APUAXgB6gC0AP0AWgB+gC0AN0AegBuvMAvwEAAP//bhC/yZZzBIIAAAAASUVORK5CYII="

		actual := palettedraster.String()

		if expected != actual {
			t.Errorf("The actual serialized value is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}
