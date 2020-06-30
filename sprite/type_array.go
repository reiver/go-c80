package c80sprite

func (receiver Type) Array() Array {

	var array Array

	copy(array[:], receiver)

	return array
}
