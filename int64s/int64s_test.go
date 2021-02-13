package int64s

import (
	"testing"
)

func TestContains(t *testing.T) {
	ints := []int64{1, 2, 3}

	if !Contains(ints, 1) {
		t.Errorf("expected contains to return true")
	}

	if Contains(ints, 4) {
		t.Errorf("expected contains to return false")
	}
}
