package path_sum

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.True(t, hasPathSum(BuildTree([]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}), 22))
	assert.False(t, hasPathSum(BuildTree([]any{1, 2, 3}), 5))
	assert.False(t, hasPathSum(BuildTree([]any{}), 0))
}
