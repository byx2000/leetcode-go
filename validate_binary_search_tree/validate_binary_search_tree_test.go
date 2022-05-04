package validate_binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test1(t *testing.T) {
	assert.True(t, isValidBST1(BuildTree([]any{2, 1, 3})))
	assert.False(t, isValidBST1(BuildTree([]any{5, 1, 4, nil, nil, 3, 6})))
}

func Test2(t *testing.T) {
	assert.True(t, isValidBST2(BuildTree([]any{2, 1, 3})))
	assert.False(t, isValidBST2(BuildTree([]any{5, 1, 4, nil, nil, 3, 6})))
}
