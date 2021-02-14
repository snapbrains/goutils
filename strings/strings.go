package strings

import (
	"crypto/sha1"
	"sort"
	str "strings"
)

// GetAdded takes two slices of strings as inputs and returns a new slice
// containing the strings that have been added
func GetAdded(original []string, modified []string) []string {
	added := []string{}
	for _, s := range modified {
		if !Contains(original, s) {
			added = append(added, s)
		}
	}
	return added
}

// GetRemoved takes two slices of strings as inputs and returns a new slice
// containing the strings that have been removed
func GetRemoved(original []string, modified []string) []string {
	removed := []string{}
	for _, s := range original {
		if !Contains(modified, s) {
			removed = append(removed, s)
		}
	}
	return removed
}

// Contains returns true if the slice contains the provided value
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// HashIt turns a slice of strings into a hash
func HashIt(s []string) string {
	sort.Strings(s)
	s = Unique(s)
	joined := str.Join(s, "")

	h := sha1.New()
	h.Write([]byte(joined))
	return string(h.Sum(nil))
}

// Unique takes a string slice and returns a new slice with duplicate values removed
func Unique(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
