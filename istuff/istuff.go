package istuff

// Include returns true if the slice includes the provided value
func Include(list []int, x int) bool {
	for _, item := range list {
		if item == x {
			return true
		}
	}
	return false
}
