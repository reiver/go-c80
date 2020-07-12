package c80textmatrix_test

import (
	"github.com/reiver/go-c80/textmatrix"

	"testing"
)

func TestType_Clear(t *testing.T) {

	var buffer [c80textmatrix.Len]byte

	var textmatrix c80textmatrix.Type = c80textmatrix.Type(buffer[:])

	textmatrix.Clear()

	{
		runes := textmatrix.Runes()

		runes[ 0] = 'H'
		runes[ 1] = 'E'
		runes[ 2] = 'L'
		runes[ 3] = 'L'
		runes[ 4] = 'O'
		runes[ 5] = ' '
		runes[ 6] = 'W'
		runes[ 7] = 'O'
		runes[ 8] = 'R'
		runes[ 9] = 'L'
		runes[10] = 'D'
		runes[11] = '!'

		runes[ 96] = 'B'
		runes[ 97] = 'e'
		runes[ 98] = 'w'
		runes[ 99] = 'a'
		runes[100] = 'r'
		runes[101] = 'e'
		runes[102] = ' '
		runes[103] = 'I'
		runes[104] = ' '
		runes[105] = 'L'
		runes[106] = 'i'
		runes[107] = 'v'
		runes[108] = 'e'
		runes[109] = '!'

		runes[len(runes)-7] = 'T'
		runes[len(runes)-6] = 'H'
		runes[len(runes)-5] = 'E'
		runes[len(runes)-4] = ' '
		runes[len(runes)-3] = 'E'
		runes[len(runes)-2] = 'N'
		runes[len(runes)-1] = 'D'
	}

	{
		runes := textmatrix.Runes()

		numNotCleared := 0

		for _,r := range runes {
			if ' ' != r {
				numNotCleared++
			}
		}

		if notExpected, actual := 0, numNotCleared; notExpected == actual {
				t.Error("The actual number of non-cleared runes is not what was expected.")
				t.Log("NOT EXPECTED:", notExpected)
				t.Log("ACTUAL:      ", actual)
			return
		}
	}

	textmatrix.Clear()

	{
		runes := textmatrix.Runes()

		for i,r := range runes {
			if expected, actual := ' ', r; expected != actual {
				t.Errorf("For rune #%d, actual value not what was expected.", i)
				t.Logf("EXPECTED: %q (%d)", expected, expected)
				t.Logf("ACTUAL:   %q (%d)", actual, actual)
				continue
			}
		}
	}
}
