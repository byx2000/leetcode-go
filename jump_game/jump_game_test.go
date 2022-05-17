package jump_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.True(t, canJump([]int{2, 3, 1, 1, 4}))
	assert.False(t, canJump([]int{3, 2, 1, 0, 4}))
	assert.False(t, canJump([]int{0, 1}))
}
