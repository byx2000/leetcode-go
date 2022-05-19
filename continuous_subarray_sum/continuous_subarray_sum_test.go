package continuous_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.True(t, checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	assert.True(t, checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	assert.False(t, checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
	assert.True(t, checkSubarraySum([]int{2, 4, 3}, 6))
}
