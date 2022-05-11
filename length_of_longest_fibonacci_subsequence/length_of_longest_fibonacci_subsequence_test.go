package length_of_longest_fibonacci_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 5, lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 3, lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18}))
}
