package int64s

// Contains returns true if the slice contains the provided value
func Contains(list []int64, x int64) bool {
	for _, item := range list {
		if item == x {
			return true
		}
	}
	return false
}
