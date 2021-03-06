package max_consecutive_ones_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 6, longestOnes1([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	assert.Equal(t, 10, longestOnes1([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	assert.Equal(t, 4, longestOnes1([]int{0, 0, 0, 1}, 4))
}

func Test2(t *testing.T) {
	assert.Equal(t, 6, longestOnes2([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	assert.Equal(t, 10, longestOnes2([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	assert.Equal(t, 4, longestOnes2([]int{0, 0, 0, 1}, 4))
}
