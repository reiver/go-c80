package c80

import (
	"fmt"
	"io"
	"os"
)

// Show will output the raster as an "IMAGE:base64-encoded-png" string to STDOUT.
func Show() {
	ShowTo(os.Stdout)
}

// ShowTo will output the raster as an "IMAGE:base64-encoded-png" string to ‘w’.
func ShowTo(w io.Writer) {
	fmt.Fprint(w, raster)
}
