package peak_index_in_a_mountain_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, peakIndexInMountainArray([]int{0, 1, 0}))
	assert.Equal(t, 1, peakIndexInMountainArray([]int{0, 2, 1, 0}))
	assert.Equal(t, 1, peakIndexInMountainArray([]int{0, 10, 5, 2}))
	assert.Equal(t, 2, peakIndexInMountainArray([]int{3, 4, 5, 1}))
	assert.Equal(t, 2, peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
	assert.Equal(t, 4, peakIndexInMountainArray([]int{40, 48, 61, 75, 100, 99, 98, 39, 30, 10}))
}
