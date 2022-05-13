package maximal_square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 4, maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
	assert.Equal(t, 1, maximalSquare([][]byte{
		{'0', '1'},
		{'1', '0'},
	}))
	assert.Equal(t, 0, maximalSquare([][]byte{
		{'0'},
	}))
}
