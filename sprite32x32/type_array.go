package c80sprite32x32

func (receiver Type) Array() Array {

	var array Array

	copy(array[:], receiver)

	return array
}
