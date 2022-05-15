package word_ladder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 5, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	assert.Equal(t, 0, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
