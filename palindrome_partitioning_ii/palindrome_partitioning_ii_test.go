package palindrome_partitioning_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, minCut("aab"))
	assert.Equal(t, 0, minCut("a"))
	assert.Equal(t, 1, minCut("ab"))
}
