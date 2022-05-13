package next_greater_node_in_linked_list

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{5, 5, 0}, nextLargerNodes(BuildList([]int{2, 1, 5})))
	assert.Equal(t, []int{7, 0, 5, 5, 0}, nextLargerNodes(BuildList([]int{2, 7, 4, 3, 5})))
}
