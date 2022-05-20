package longest_palindromic_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 4, longestPalindromeSubseq("bbbab"))
	assert.Equal(t, 2, longestPalindromeSubseq("cbbd"))
	assert.Equal(t, 1, longestPalindromeSubseq("abcdef"))
}
