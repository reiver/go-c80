package c80frame

import (
	"bytes"
	"encoding/base64"
	"strings"
)

func (receiver *Type) String() string {
	var buffer strings.Builder

	{
		buffer.WriteString("IMAGE:")
	}

	{
		var pngBuffer bytes.Buffer

		receiver.PNGTo(&pngBuffer)

		base64Encoded := base64.StdEncoding.EncodeToString(pngBuffer.Bytes())
		buffer.WriteString(base64Encoded)
	}

	return buffer.String()
}
