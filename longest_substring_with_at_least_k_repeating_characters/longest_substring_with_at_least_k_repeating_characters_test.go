package longest_substring_with_at_least_k_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, longestSubstring("aaabb", 3))
	assert.Equal(t, 5, longestSubstring("ababbc", 2))
}
