package route_between_nodes_lcci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.True(t, findWhetherExistsPath(3, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}}, 0, 2))
	assert.True(t, findWhetherExistsPath(5, [][]int{{0, 1}, {0, 2}, {0, 4}, {0, 4}, {0, 1}, {1, 3}, {1, 4}, {1, 3}, {2, 3}, {3, 4}}, 0, 4))
}
