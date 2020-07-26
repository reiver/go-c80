package c80

// Memory returns the entire memory of the c80 machine.
//
// Reading and writing to this give you access to everything.
func Memory() []uint8 {
	return machine.Memory()
}
