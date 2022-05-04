package kth_smallest_element_in_a_bst

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, kthSmallest(BuildTree([]any{3, 1, 4, nil, 2}), 1))
	assert.Equal(t, 3, kthSmallest(BuildTree([]any{5, 3, 6, 2, 4, nil, nil, 1}), 3))
}
