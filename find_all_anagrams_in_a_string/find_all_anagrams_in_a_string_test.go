package find_all_anagrams_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{}, findAnagrams("aaa", "aaaa"))
	assert.Equal(t, []int{0, 6}, findAnagrams("cbaebabacd", "abc"))
	assert.Equal(t, []int{0, 1, 2}, findAnagrams("abab", "ab"))
}
