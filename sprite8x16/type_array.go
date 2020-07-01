package c80sprite8x16

func (receiver Type) Array() Array {

	var array Array

	copy(array[:], receiver)

	return array
}
