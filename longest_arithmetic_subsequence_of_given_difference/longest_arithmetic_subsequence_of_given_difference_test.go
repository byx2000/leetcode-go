package longest_arithmetic_subsequence_of_given_difference

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 4, longestSubsequence1([]int{1, 2, 3, 4}, 1))
	assert.Equal(t, 1, longestSubsequence1([]int{1, 3, 5, 7}, 1))
	assert.Equal(t, 4, longestSubsequence1([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}

func Test2(t *testing.T) {
	assert.Equal(t, 4, longestSubsequence2([]int{1, 2, 3, 4}, 1))
	assert.Equal(t, 1, longestSubsequence2([]int{1, 3, 5, 7}, 1))
	assert.Equal(t, 4, longestSubsequence2([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}
