package detect_cycles_in_2d_grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.True(t, containsCycle([][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'a', 'a', 'a'},
	}))
	assert.True(t, containsCycle([][]byte{
		{'c', 'c', 'c', 'a'},
		{'c', 'd', 'c', 'c'},
		{'c', 'c', 'e', 'c'},
		{'f', 'c', 'c', 'c'},
	}))
	assert.False(t, containsCycle([][]byte{
		{'a', 'b', 'b'},
		{'b', 'z', 'b'},
		{'b', 'b', 'a'},
	}))
	assert.False(t, containsCycle([][]byte{
		{'b', 'a', 'c'},
		{'c', 'a', 'c'},
		{'d', 'd', 'c'},
		{'b', 'c', 'c'},
	}))
}
