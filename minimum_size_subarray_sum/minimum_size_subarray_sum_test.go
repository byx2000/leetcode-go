package minimum_size_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 2, minSubArrayLen1(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 1, minSubArrayLen1(4, []int{1, 4, 4}))
	assert.Equal(t, 0, minSubArrayLen1(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 2, minSubArrayLen2(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 1, minSubArrayLen2(4, []int{1, 4, 4}))
	assert.Equal(t, 0, minSubArrayLen2(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
