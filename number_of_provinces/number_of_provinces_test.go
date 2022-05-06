package number_of_provinces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))
	assert.Equal(t, 3, findCircleNum([][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}))
}
