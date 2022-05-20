package contiguous_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, findMaxLength([]int{0, 1}))
	assert.Equal(t, 2, findMaxLength([]int{0, 1, 0}))
}
