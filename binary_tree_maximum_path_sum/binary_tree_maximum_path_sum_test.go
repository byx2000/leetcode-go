package binary_tree_maximum_path_sum

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, maxPathSum(BuildTree([]any{1, 2, 3})), 6)
	assert.Equal(t, maxPathSum(BuildTree([]any{-10, 9, 20, nil, nil, 15, 7})), 42)
	assert.Equal(t, maxPathSum(BuildTree([]any{2, -1})), 2)
}
