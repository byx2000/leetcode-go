package word_break_ii

import (
	"github.com/stretchr/testify/assert"
	"leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, common.SliceToSet([]string{"cats and dog", "cat sand dog"}), common.SliceToSet(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})))
	assert.Equal(t, common.SliceToSet([]string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}), common.SliceToSet(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})))
	assert.Equal(t, common.SliceToSet([]string{}), common.SliceToSet(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})))
}
