package wildcard_matching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.False(t, isMatch("aa", "a"))
	assert.True(t, isMatch("aa", "*"))
	assert.False(t, isMatch("cb", "?a"))
	assert.True(t, isMatch("adceb", "*a*b"))
	assert.False(t, isMatch("acdcb", "a*c?b"))
	assert.True(t, isMatch("", "******"))
	assert.False(t, isMatch("mississippi", "m??*ss*?i*pi"))
}
