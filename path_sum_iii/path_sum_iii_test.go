package path_sum_iii

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, pathSum(BuildTree([]any{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}), 8))
	assert.Equal(t, 3, pathSum(BuildTree([]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}), 22))
}
