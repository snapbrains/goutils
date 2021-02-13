package int64s

// ContainsInt64 returns true if the slice contains the provided value
func ContainsInt64(s []int64, str int64) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
