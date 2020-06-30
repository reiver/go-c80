package c80

import (
	"fmt"
	"io"
	"os"
)

// Show will output the raster as an "IMAGE:base64-encoded-png" string to STDOUT.
//
// (Where "base64-encoded-png" is replaced with a base64 encoding of the PNG encoding
// of the raster image.)
//
// If you are using the Go Playground — https://play.golang.org/ — then
// you could run:…
//
//	c80.Show()
//
// … at the end of your program to be able to see the raster image.
//
// If the only thing you output in the Go Playground — https://play.golang.org/ — is
// "IMAGE:base64-encoded-png", then the Go Playground — https://play.golang.org/ — will
// display the actual image (rather than the "IMAGE:base64-encoded-png" text).
//
// For example, if this code:…
//
//	c80.Dye(3)
//	c80.Rectangle(40,60, 64,48 4)
//
//	c80.Show()
//
// … is run in the Go Playground — https://play.golang.org/ — then the Go Playground
// will show a 128×192 image with a rectangle of one color, over a background of another color.
func Show() {
	ShowTo(os.Stdout)
}

// ShowTo will output the raster as an "IMAGE:base64-encoded-png" string to ‘w’.
//
// (Where "base64-encoded-png" is replaced with a base64 encoding of the PNG encoding
// of the raster image.)
//
// ShowTo is similar to Show, except it lets you choose where you want to output the
// "IMAGE:base64-encoded-png" string (using ‘w’).
func ShowTo(w io.Writer) {
	fmt.Fprint(w, machine.Raster(0))
}
