package replace_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, "the cat was rat by the bat", replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	assert.Equal(t, "a a b c", replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}
