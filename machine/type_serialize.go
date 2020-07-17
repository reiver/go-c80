package c80machine

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
)

func (receiver *Type) Serialize() string {
	if nil == receiver {
		return "ERROR: nil receiver"
	}

	var pngBuffer bytes.Buffer

	if err := receiver.PNGTo(&pngBuffer); nil != err {
		return fmt.Sprintf("ERROR: could not write PNG data: %s", err)
	}

	var buffer strings.Builder
	{
		buffer.WriteString("IMAGE:")

		base64Encoded := base64.StdEncoding.EncodeToString(pngBuffer.Bytes())

		buffer.WriteString(base64Encoded)
	}

	return buffer.String()
}
