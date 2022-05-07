package partition_equal_subset_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.True(t, canPartition([]int{1, 5, 11, 5}))
	assert.False(t, canPartition([]int{1, 2, 3, 5}))
}
