package c80_test

import (
	"github.com/reiver/go-c80"

	"github.com/reiver/go-sprite8x8"

	"fmt"
)

func ExampleSprite() {

	// Set the color palette to "vt".
	err := c80.Colorize("vt")
	if nil != err {
		fmt.Printf("ERROR: could not set the palette: %s", err)
		return
	}

	// These index values are specific to the "vt" color palette.
	const black = 0
	const blue  = 4
	const white = 7



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
			black, black, black, black, black, black, black, black,
			black, black, black, black, black, black, black, black,
			black, black, black, black, black, black, black, black,
			black, black, black, black, black, black, black, black,
			black, black, black, black, black, black, black, black,
			black, black, black, black, black, white, white, white,
			white, white, white, white, white, white, black, white,
			white, black, white, black, black, white, white, white,
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
	// IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAIAAADUvDoHAAADhUlEQVR4nOzXsY3iQBSAYbxySTtFXDO7vVwLFAE9uBV64AJLDg5ij8T/fYE1TPSC98t4XZbl8s7z5/b2Hj7J1+wBYCYBkCYA0gRA2vrf723b9sP39TFjHjjVemz8YYzxegkf6Wvf+P25H2w/Hcvl9z57BpjGRzBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBtnT1AzvL3z9v758/t9FnwBqBNAKQJgDQBkOYjeLJt2/bD9/Uxe5YiAZzt2PjDGOP1knP4CzTBGGN/7gfbP9Fy+b3PngGm8QYgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZD2LwAA///KYCVGxq8sfwAAAABJRU5ErkJggg==
}
