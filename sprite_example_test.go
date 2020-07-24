package c80_test

import (
	"github.com/reiver/go-c80"

	"github.com/reiver/go-sprite8x8"

	"fmt"
	"image/color"
)

func ExampleSprite() {

	// Set the color palette to "vt".
	err := c80.Colorize("vt")
	if nil != err {
		fmt.Printf("ERROR: could not set the palette: %s", err)
		return
	}

	// Get the closest colors in the current palette to these colors.
	var blue     uint8 = c80.ColorIndex(color.RGBA{0,0,255,0})
	var white    uint8 = c80.ColorIndex(color.RGBA{255,255,255,255})
	var trnsprnt uint8 = c80.ColorIndex(color.RGBA{0,0,0,0})



	// Clear.
	err = c80.Draw(c80.Dye(blue))
	if nil != err {
		fmt.Printf("ERROR: could not clear the frame: %s", err)
		return
	}



	// The id of the sprite.
	//
	// We could have picked any number from 0 to 255.
	// In this example, we simple just chose this one.
	const id = 200

	// Here we define what the sprite will look like.
	//
	// This is an 8×8 grid.
	//
	// This sprite art — which nowadays many call: pixel-art — is of a key,
	// that looks something like:
	//
	//	     ▄▄▄
	//	█▀█▀▀█▄█
	{
		var buffer [8*8]uint8 = [8*8]uint8{
			trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt,
			trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt,
			trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt,
			trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt,
			trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt,
			trnsprnt, trnsprnt, trnsprnt, trnsprnt, trnsprnt, white,    white,    white,
			white,    white,    white,    white,    white,    white,    trnsprnt, white,
			white,    trnsprnt, white,    trnsprnt, trnsprnt, white,    white,    white,
		}

		sprite := c80.Sprite("8x8", id).(sprite8x8.Paletted)

		p := sprite.Pix

		copy(p, buffer[:])

	}



	// We are going to draw that sprite twice.
	//
	// This is the 1st time we draw the sprite.
	//
	// We don't relocate the position of the sprite this time.
	// So the sprite will just appear in the top-left corner.
	c80.Draw(c80.Sprite("8x8", id))



	// This is the (x,y) coordinate of where we will draw the sprite
	// the 2nd time.
	//
	// This (x,y) location corresponds to the sprite's (new) top-left corner.
	const x,y = 123,139

	// This is the 2nd time we draw the (same) sprite.
	//
	// Except this time we have relocated it to (x,y)=(123,139) (rather than
	// (x,y)=(0,0) like the 1st time we drew it).
	//
	// Here (x,y)=(123,139) corresponds to the sprite's top-left corner.
	c80.Draw(c80.Relocate(x,y, c80.Sprite("8x8", id)))


	// Here we draw a sprite at (x,y)=(200,180)
	c80.Draw(c80.Relocate(200,180, c80.Sprite("8x8", 247)))


	// Show the raster image.
	//
	// This will output something like:
	//
	//	IMAGE:base64-encoded-png
	//
	// Where "base64-encoded-png" is replaced with the base64 encoding of a PNG version of the raster image.
	serialized := c80.Serialize()
	fmt.Print(serialized)

	// Output:
	// IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAIAAADUvDoHAAADvUlEQVR4nOzdTW7bMBgAUX2Cb9Tc/wTplaKisSu0ibeyhM57K1orIuDoh1nwtizbAlXr2ROAMwmANAGQJgDSbl9+b38+iWdOmA282G37tgk0s3y/CP+ldb/ZzzwGVj8d4/8AlPkIJk0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEi7nT2Bum17DGZOnkmTAF5tX/G7mScXeQ2vQCe43+xnHgOr/0SzLP78dHkCkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDRHJHGgbXmcfDZXPYfFE4Cj7Kv/y/hSBMCB1s8DANeZ9aqHYDojjKPc7/o/39/vP9/efpw9oycEwCHuq3+/8X98HoV5wS8BH8Ec5ffbzzb/jK/3HuQJwFFm1s9TkD/2wdkzekIAHOLvDdArb4YKgDTboKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkPYrAAD//99QOU2R1Sq/AAAAAElFTkSuQmCC
}
