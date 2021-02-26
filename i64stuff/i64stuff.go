package i64stuff

// Include returns true if the slice includes the provided value
func Include(list []int64, x int64) bool {
	for _, item := range list {
		if item == x {
			return true
		}
	}
	return false
}
