package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 0, lengthOfLongestSubstring1(""))
	assert.Equal(t, 1, lengthOfLongestSubstring1("a"))
	assert.Equal(t, 3, lengthOfLongestSubstring1("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring1("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring1("pwwkew"))
}

func Test2(t *testing.T) {
	assert.Equal(t, 0, lengthOfLongestSubstring2(""))
	assert.Equal(t, 1, lengthOfLongestSubstring2("a"))
	assert.Equal(t, 3, lengthOfLongestSubstring2("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring2("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring2("pwwkew"))
}
