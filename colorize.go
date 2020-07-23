package c80

// Colorize sets the color palette.
//
// Example
//
// Here are some examples of Colorize being used:
//
//	err := c80.Colorize("tia")
func Colorize(a ...interface{}) error {
	return machine.Colorize(a...)
}
