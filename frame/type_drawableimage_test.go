package c80frame_test

import (
	"github.com/reiver/go-font8x8"

	"github.com/reiver/go-c80/frame"

	"image"
	"image/color"
	"image/draw"

	"testing"
)

func TestType_DrawableImage(t *testing.T) {

	var frame c80frame.Type

	var bytes []byte
	{
		bytes = frame.Bytes()

		if length := len(bytes); 0 >= length {
			t.Errorf("The actual number of bytes is less than expected.")
			t.Logf("LENGTH: %d", length)
			return
		}
		if expected, actual := c80frame.ByteSize, len(bytes); expected != actual {
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
		drawable := frame.DrawableImage()

		c := color.NRGBA{
			R:bg1.Red,
			G:bg1.Green,
			B:bg1.Blue,
			A:bg1.Alpha,
		}

		for y:=0; y<c80frame.Height; y++ {
			for x:=0; x<c80frame.Width; x++ {
				drawable.Set(x,y, c)
			}
		}
	}

	// Check to see what is there AFTER the 1st time we draw anything is what we expect.
	{
		image := frame.DrawableImage()

		for y:=0; y<c80frame.Height; y++ {
			for x:=0; x<c80frame.Width; x++ {
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
					t.Log(frame.String())
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

		drawableImage := frame.DrawableImage()

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
	for offset, value := range bytes {
		channel := offset % 4

		var expected byte
		switch channel {
		case 0:
			expected = bg2.Red
		case 1:
			expected = bg2.Green
		case 2:
			expected = bg2.Blue
		case 3:
			expected = bg2.Alpha
		default:
			t.Error("This should never happen.")
			return
		}

		if actual := value; expected != actual {
			t.Errorf("For offset=%d (and thus channel=%d), the actual value for the pixel color channel is not what was expected.", offset, channel)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Log(frame.String())
			return
		}
	}

	const fontWidth  = 8
	const fontHeight = 8

	const printX = 10
	const printY = 20

	// Draw the letter "N".
	{
		drawable := frame.DrawableImage()

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
	for offset, value := range bytes {
		x := (offset / 4) % c80frame.Width
		y := (offset / 4) / c80frame.Width

		// If we are on the stop where we drew the character.
		if (printY <= y && y < (printY+fontHeight)) && (printX <= x && x < (printX+fontWidth)) {
			channel := offset % 4

			switch channel {
			case 0,1,2:
				if expected1, expected2, actual := byte(0x00), byte(0xff), value; expected1 != actual && expected2 != actual {
					t.Errorf("For offset=%d (and thus channel=%d), the actual value for the pixel color channel is not what was expected.", offset, channel)
					t.Log("WE ARE IN THE AREA WE DREW THE CHARACTER....")
					t.Logf("EXPECTED 1: %d", expected1)
					t.Logf("EXPECTED 2: %d", expected2)
					t.Logf("ACTUAL:     %d", actual)
					t.Log(frame.String())
					return
				}
			case 3:
				if expected, actual := byte(0xff), value; expected != actual {
					t.Errorf("For offset=%d (and thus channel=%d), the actual value for the pixel color channel is not what was expected.", offset, channel)
					t.Log("WE ARE IN THE AREA WE DREW THE CHARACTER....")
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					t.Log(frame.String())
					return
				}
			default:
				t.Error("This should never happen.")
				return
			}
			continue
		}

		channel := offset % 4

		var expected byte
		switch channel {
		case 0:
			expected = bg2.Red
		case 1:
			expected = bg2.Green
		case 2:
			expected = bg2.Blue
		case 3:
			expected = bg2.Alpha
		default:
			t.Error("This should never happen.")
			return
		}

		if actual := value; expected != actual {
			t.Errorf("For offset=%d (and thus channel=%d), the actual value for the pixel color channel is not what was expected.", offset, channel)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Log(frame.String())
			return
		}
	}

	{
		expected := "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAIAAADUvDoHAAADWUlEQVR4nOzbIW6FQBRAUWh+uqbu33RNVVRMUlHDfEXy7jlugkFwmeElvK7vzwOqPp6+AXiSAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGmvm+tfP8dxnOe5Vtd1raU/iZlhawdYzz3Ms3sE0gAjbQWwjkAaYJ7dHeDvMwAmeWMKpAHmMQYl7WYM+u+tbxNgmJsAzPuZzRGINAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSfgMAAP//5bwU5weyrzMAAAAASUVORK5CYII="

		actual := frame.String()

		if expected != actual {
			t.Errorf("The actual serialized value is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}
