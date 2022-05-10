package longest_increasing_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	assert.Equal(t, 1, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
