package maximum_length_of_repeated_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	assert.Equal(t, 5, findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))

}
