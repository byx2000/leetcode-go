package subarray_sums_divisible_by_k

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 7, subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	assert.Equal(t, 0, subarraysDivByK([]int{5}, 9))
	assert.Equal(t, 2, subarraysDivByK([]int{-1, 2, 9}, 2))
	assert.Equal(t, 2, subarraysDivByK([]int{2, -2, 2, -4}, 6))
}
