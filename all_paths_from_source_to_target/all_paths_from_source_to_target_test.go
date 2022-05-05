package all_paths_from_source_to_target

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, [][]int{{0, 1, 3}, {0, 2, 3}}, allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	assert.Equal(t, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}, allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}
