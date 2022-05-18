package shortest_bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, shortestBridge([][]int{
		{0, 1},
		{1, 0},
	}))
	assert.Equal(t, 2, shortestBridge([][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 1},
	}))
	assert.Equal(t, 1, shortestBridge([][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}))
}
