package next_greater_element_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{2, -1, 2}, nextGreaterElements([]int{1, 2, 1}))
	assert.Equal(t, []int{2, 3, 4, -1, 4}, nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
