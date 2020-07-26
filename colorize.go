package c80

// Colorize sets the color palette.
//
// Example
//
// Here are some examples of Colorize being used:
//
//	err := c80.Colorize("tia")
//
// Valid color palettes include:
//
// • "gb"
//
// • "greer"
//
// • "gruvbox"
//
// • "html3"
//
// • "lospec500"
//
// • "nes"
//
// • "rkbv"
//
// • "solarized"
//
// • "tia"
func Colorize(a ...interface{}) error {
	return machine.Colorize(a...)
}
