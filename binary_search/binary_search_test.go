package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
