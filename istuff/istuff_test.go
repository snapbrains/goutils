package istuff

import (
	"testing"
)

func TestInclude(t *testing.T) {
	ints := []int{1, 2, 3}

	if !Include(ints, 1) {
		t.Errorf("expected include to return true")
	}

	if Include(ints, 4) {
		t.Errorf("expected include to return false")
	}
}
