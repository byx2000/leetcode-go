package nearest_exit_from_entrance_in_maze

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, nearestExit([][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'},
	}, []int{1, 2}))
	assert.Equal(t, 2, nearestExit([][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'},
	}, []int{1, 0}))
	assert.Equal(t, -1, nearestExit([][]byte{
		{'.', '+'},
	}, []int{0, 0}))
}
