package minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "a", minWindow("a", "a"))
	assert.Equal(t, "", minWindow("a", "aa"))
}
