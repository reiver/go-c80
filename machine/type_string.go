package c80machine

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"strings"
)

func (receiver Type) String() string {
	var buffer strings.Builder

	{
		buffer.WriteString("IMAGE:")
	}

	{
		var pngBuffer bytes.Buffer

		png.Encode(&pngBuffer, receiver)

		base64Encoded := base64.StdEncoding.EncodeToString(pngBuffer.Bytes())
		buffer.WriteString(base64Encoded)
	}

	return buffer.String()
}