package find_the_longest_substring_containing_vowels_in_even_counts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 13, findTheLongestSubstring("eleetminicoworoep"))
	assert.Equal(t, 5, findTheLongestSubstring("leetcodeisgreat"))
	assert.Equal(t, 6, findTheLongestSubstring("bcbcbc"))
}
