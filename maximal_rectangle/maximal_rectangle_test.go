package maximal_rectangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 6, maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
	assert.Equal(t, 0, maximalRectangle([][]byte{}))
	assert.Equal(t, 0, maximalRectangle([][]byte{{'0'}}))
	assert.Equal(t, 1, maximalRectangle([][]byte{{'1'}}))
	assert.Equal(t, 0, maximalRectangle([][]byte{{'0', '0'}}))
}
