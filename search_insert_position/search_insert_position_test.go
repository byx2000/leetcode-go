package search_insert_position

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
}
