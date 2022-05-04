package longest_univalue_path

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, longestUnivaluePath(BuildTree([]any{5, 4, 5, 1, 1, 5})), 2)
	assert.Equal(t, longestUnivaluePath(BuildTree([]any{1, 4, 5, 4, 4, 5})), 2)
}
