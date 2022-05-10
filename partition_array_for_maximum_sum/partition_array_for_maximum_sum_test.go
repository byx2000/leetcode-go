package partition_array_for_maximum_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 84, maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
	assert.Equal(t, 83, maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4))
	assert.Equal(t, 1, maxSumAfterPartitioning([]int{1}, 1))
}
