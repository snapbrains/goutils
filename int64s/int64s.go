package int64s

// Contains returns true if the slice contains the provided value
func Contains(s []int64, str int64) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
