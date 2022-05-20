package palindrome_partitioning

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, [][]string{{"a", "a", "b"}, {"aa", "b"}}, partition("aab"))
	assert.Equal(t, [][]string{{"g", "o", "o", "g", "l", "e"}, {"g", "oo", "g", "l", "e"}, {"goog", "l", "e"}}, partition("google"))
	assert.Equal(t, [][]string{{"a"}}, partition("a"))
}
