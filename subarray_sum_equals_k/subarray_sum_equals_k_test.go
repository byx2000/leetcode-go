package subarray_sum_equals_k

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 2, subarraySum1([]int{1, 1, 1}, 2))
	assert.Equal(t, 2, subarraySum1([]int{1, 2, 3}, 3))
	assert.Equal(t, 4, subarraySum1([]int{1, 2, 1, 2, 1}, 3))
}

func Test2(t *testing.T) {
	assert.Equal(t, 2, subarraySum2([]int{1, 1, 1}, 2))
	assert.Equal(t, 2, subarraySum2([]int{1, 2, 3}, 3))
	assert.Equal(t, 4, subarraySum2([]int{1, 2, 1, 2, 1}, 3))
}
