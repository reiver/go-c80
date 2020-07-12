package c80textmatrix_test

import (
	"github.com/reiver/go-c80/textmatrix"

	"testing"
)

func TestType_Runes(t *testing.T) {

	var buffer [c80textmatrix.Len]byte

	var textmatrix c80textmatrix.Type = c80textmatrix.Type(buffer[:])

	runes := textmatrix.Runes()

	if expected, actual := len(buffer)/4, len(runes); expected != actual {
		t.Errorf("The actual len(runes) is not what was expected.")
		t.Log("EXPECTED:", expected)
		t.Log("ACTUAL:  ", actual)
		return
	}

	if expected, actual := cap(buffer)/4, cap(runes); expected != actual {
		t.Errorf("The actual len(runes) is not what was expected.")
		t.Log("EXPECTED:", expected)
		t.Log("ACTUAL:  ", actual)
		return
	}

	runes[ 0] = 'A'
	runes[ 1] = 'B'
	runes[ 2] = 'C'
	runes[ 3] = 'D'
	runes[ 4] = 'E'
	runes[ 5] = 'F'
	runes[ 6] = 'G'
	runes[ 7] = 'H'
	runes[ 8] = 'I'
	runes[ 9] = 'J'
	runes[10] = 'K'
	runes[11] = 'L'
	runes[12] = 'M'
	runes[13] = 'N'
	runes[14] = 'O'
	runes[15] = 'P'
	runes[16] = 'Q'
	runes[17] = 'R'
	runes[18] = 'S'
	runes[19] = 'T'
	runes[20] = 'U'
	runes[21] = 'V'
	runes[22] = 'W'
	runes[23] = 'X'
	runes[24] = 'Y'
	runes[25] = 'Z'

	for i:=0; i<26; i++ {
		p0 := i*4
		p1 := i*4 + 1
		p2 := i*4 + 2
		p3 := i*4 + 3

		var c byte = byte(i) + byte('A')

		if actual, expected := buffer[p0], c; expected != actual {
			t.Errorf("For i=%d, c=%q, and byte 0, actual value is not what was expected.", i, string(rune(c)))
			t.Logf("EXPECTED: %d (%q)", expected, string(rune(expected)))
			t.Logf("ACTUAL:   %d (%q)", actual, string(rune(actual)))
			continue
		}
		if actual, expected := buffer[p1], byte(0); expected != actual {
			t.Errorf("For i=%d, c=%q, and byte 1, actual value is not what was expected.", i, string(rune(c)))
			t.Logf("EXPECTED: %d (%q)", expected, string(rune(expected)))
			t.Logf("ACTUAL:   %d (%q)", actual, string(rune(actual)))
			continue
		}
		if actual, expected := buffer[p2], byte(0); expected != actual {
			t.Errorf("For i=%d, c=%q, and byte 2, actual value is not what was expected.", i, string(rune(c)))
			t.Logf("EXPECTED: %d (%q)", expected, string(rune(expected)))
			t.Logf("ACTUAL:   %d (%q)", actual, string(rune(actual)))
			continue
		}
		if actual, expected := buffer[p3], byte(0); expected != actual {
			t.Errorf("For i=%d, c=%q, and byte 3, actual value is not what was expected.", i, string(rune(c)))
			t.Logf("EXPECTED: %d (%q)", expected, string(rune(expected)))
			t.Logf("ACTUAL:   %d (%q)", actual, string(rune(actual)))
			continue
		}
	}

	for i:=27*4; i<len(buffer); i++ {
		if actual, expected := buffer[i], byte(0); expected != actual {
			t.Errorf("For i=%d, actual value is not what was expected.", i)
			t.Logf("EXPECTED: %d (%q)", expected, string(rune(expected)))
			t.Logf("ACTUAL:   %d (%q)", actual, string(rune(actual)))
			continue
		}
	}
}
