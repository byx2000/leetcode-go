package house_robber_iii

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 7, rob1(BuildTree([]any{3, 2, 3, nil, 3, nil, 1})))
	assert.Equal(t, 9, rob1(BuildTree([]any{3, 4, 5, 1, 3, nil, 1})))
}

func Test2(t *testing.T) {
	assert.Equal(t, 7, rob2(BuildTree([]any{3, 2, 3, nil, 3, nil, 1})))
	assert.Equal(t, 9, rob2(BuildTree([]any{3, 4, 5, 1, 3, nil, 1})))
}
