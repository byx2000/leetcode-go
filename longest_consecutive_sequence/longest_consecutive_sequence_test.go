package longest_consecutive_sequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 0, longestConsecutive([]int{}))
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	assert.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
