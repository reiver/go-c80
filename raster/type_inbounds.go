package c80raster

func (receiver Type) InBounds(x int, y int) bool {
	if Width <= x {
		return false
	}
	if Height <= y {
		return false
	}

	if x < 0 {
		return false
	}
	if y < 0 {
		return false
	}

	return true
}
