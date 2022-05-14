package redundant_connection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{2, 3}, findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	assert.Equal(t, []int{1, 4}, findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
}
