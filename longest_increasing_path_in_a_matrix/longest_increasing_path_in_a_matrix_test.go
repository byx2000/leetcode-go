package longest_increasing_path_in_a_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 4, longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}))
	assert.Equal(t, 4, longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}))
	assert.Equal(t, 1, longestIncreasingPath([][]int{
		{1},
	}))
	assert.Equal(t, 2, longestIncreasingPath([][]int{
		{1, 2},
	}))
}
