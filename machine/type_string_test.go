package c80machine_test

import (
	"github.com/reiver/go-c80/machine"

	"testing"
)

func TestTypeString(t *testing.T) {

	tests := []struct{
		Machine *c80machine.Type
		Expected string
	}{
		{
			Machine: func()*c80machine.Type{
				var machine c80machine.Type

				return &machine
			}(),
			Expected: "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAH8AAAC/EAYAAACHxd/kAAACKElEQVR4nOzTAREAAAgDIc/+oRfkoQN/QI74ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgSJD0HiQ5D4ECQ+BIkPQeJDkPgQJD4EiQ9B4kOQ+BAkPgQtAAD//y2LAX+xz27zAAAAAElFTkSuQmCC",
		},
		{
			Machine: func()*c80machine.Type{
				var machine c80machine.Type

				const index = 2

				machine.Palette().Color(index).Poke(0,0,0,255)
				machine.Raster(0).Dye(index)

				return &machine
			}(),
			Expected: "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAH8AAAC/EAIAAAAIp0izAAAB+UlEQVR4nOzSAREAAAgDIc/+oRfkIQN/kKQ+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hOlPlHqE6U+UeoTpT5R6hO1AAAA//9GKAF/GTWbmAAAAABJRU5ErkJggg==",
		},
		{
			Machine: func()*c80machine.Type{
				var machine c80machine.Type

				const index = 1

				machine.Palette().Color(index).Poke(255,199,6,255)
				machine.Raster(0).Dye(index)

				return &machine
			}(),
			Expected: "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAH8AAAC/EAIAAAAIp0izAAACHUlEQVR4nOzSQQ2AQBDAQELI+Xd1shYZ++iMgj76zdx7zgMx73YA7LA+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xNlfaKsT5T1ibI+UdYnyvpEWZ8o6xP1BwAA///4GAUZT5N97AAAAABJRU5ErkJggg==",
		},
		{
			Machine: func()*c80machine.Type{
				var machine c80machine.Type

				const index = 11

				machine.Palette().Color(index).Poke(0,111,184,255)
				machine.Raster(0).Dye(index)

				return &machine
			}(),
			Expected: "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAH8AAAC/EAIAAAAIp0izAAACHElEQVR4nOzSMRHAMAzAQF+vfEMv8ALDg/4RaNA/c869AzHfdgDssD5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE2V9oqxPlPWJsj5R1ifK+kRZnyjrE/UCAAD///fjA8+/tABGAAAAAElFTkSuQmCC",
		},
	}

	for testNumber, test := range tests {

		actual := test.Machine.String()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value was not what was expected." , testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
