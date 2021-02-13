package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdded(t *testing.T) {
	original := []string{"red", "green", "blue"}
	modified := []string{"red", "blue", "yellow"}

	added := GetAdded(original, modified)
	assert.Equal(t, added, []string{"yellow"})
}

func TestGetRemoved(t *testing.T) {
	original := []string{"red", "green", "blue"}
	modified := []string{"red", "blue", "yellow"}

	removed := GetRemoved(original, modified)
	assert.Equal(t, removed, []string{"green"})
}
