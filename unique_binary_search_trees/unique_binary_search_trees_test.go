package unique_binary_search_trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, numTrees(1))
	assert.Equal(t, 5, numTrees(3))
	assert.Equal(t, 16796, numTrees(10))
}
