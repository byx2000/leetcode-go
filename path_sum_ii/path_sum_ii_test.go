package path_sum_ii

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}, pathSum(BuildTree([]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}), 22))
	var empty [][]int
	assert.Equal(t, empty, pathSum(BuildTree([]any{1, 2, 3}), 5))
	assert.Equal(t, empty, pathSum(BuildTree([]any{1, 2}), 0))
}
