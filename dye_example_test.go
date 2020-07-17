package c80_test

import (
	"github.com/reiver/go-c80"

	"fmt"
)

func ExampleDye() {

	// Make the entire raster image the color of color 3 in the (color) palette.
	c80.Dye(3)

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
	// IMAGE:iVBORw0KGgoAAAANSUhEUgAAAH8AAAC/BAMAAAAlex6UAAAAMFBMVEUBAQHeOCs5tUr/xwYAb7h2JnEstenMzMyAgID/AAAA/wD//wAAAP//AP8A//////9ssw/GAAAATElEQVR4nOzMQQ0AAAgEICO52b+bIe53gwDMZVYgEAgEAoFAIBAIBAKBQCAQCAQCgUAgEAgEAoFAIBAIBAKBQCAQCAQCQXvwAQAA///AzoGLUCkA2wAAAABJRU5ErkJggg==
}
