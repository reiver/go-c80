package c80sprite8x8

func (receiver Type) Array() Array {

	var array Array

	copy(array[:], receiver)

	return array
}
