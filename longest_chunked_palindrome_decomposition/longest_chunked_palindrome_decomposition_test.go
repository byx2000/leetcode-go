package longest_chunked_palindrome_decomposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 7, longestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
	assert.Equal(t, 1, longestDecomposition("merchant"))
	assert.Equal(t, 11, longestDecomposition("antaprezatepzapreanta"))
	assert.Equal(t, 3, longestDecomposition("aaa"))
	assert.Equal(t, 2, longestDecomposition("elvtoelvto"))
}
