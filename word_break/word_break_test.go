package word_break

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.True(t, wordBreak("leetcode", []string{"leet", "code"}))
	assert.True(t, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.False(t, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	assert.True(t, wordBreak("ab", []string{"a", "b"}))
}
