package number_of_operations_to_make_network_connected

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	assert.Equal(t, 2, makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
	assert.Equal(t, -1, makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}))
	assert.Equal(t, 0, makeConnected(5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}))
}
