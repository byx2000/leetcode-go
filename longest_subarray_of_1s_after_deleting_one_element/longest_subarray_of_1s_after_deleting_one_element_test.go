package longest_subarray_of_1s_after_deleting_one_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, longestSubarray([]int{1, 1, 0, 1}))
	assert.Equal(t, 5, longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	assert.Equal(t, 2, longestSubarray([]int{1, 1, 1}))
}
