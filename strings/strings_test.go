package strings

import (
	"reflect"
	"testing"
)

func TestGetAdded(t *testing.T) {
	original := []string{"red", "green", "blue"}
	modified := []string{"red", "blue", "yellow"}

	added := GetAdded(original, modified)

	if !reflect.DeepEqual(added, []string{"yellow"}) {
		t.Errorf("expected removed to contain 'yellow' but it was %v", added)
	}
}

func TestGetRemoved(t *testing.T) {
	original := []string{"red", "green", "blue"}
	modified := []string{"red", "blue", "yellow"}

	removed := GetRemoved(original, modified)
	if !reflect.DeepEqual(removed, []string{"green"}) {
		t.Errorf("expected removed to contain 'green' but it was %v", removed)
	}
}

func TestContains(t *testing.T) {
	strings := []string{"red", "green", "blue"}

	if !Contains(strings, "red") {
		t.Errorf("expected contains to return true")
	}

	if Contains(strings, "purple") {
		t.Errorf("expected contains to return false")
	}
}
