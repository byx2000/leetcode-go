package house_robber_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 0, rob([]int{0}))
	assert.Equal(t, 3, rob([]int{2, 3, 2}))
	assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 3, rob([]int{1, 2, 3}))
}
