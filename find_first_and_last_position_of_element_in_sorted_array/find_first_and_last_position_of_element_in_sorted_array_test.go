package find_first_and_last_position_of_element_in_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
}
