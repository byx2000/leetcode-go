package scramble_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.True(t, isScramble1("great", "rgeat"))
	assert.False(t, isScramble1("abcde", "caebd"))
	assert.True(t, isScramble1("a", "a"))
}

func Test2(t *testing.T) {
	assert.True(t, isScramble2("great", "rgeat"))
	assert.False(t, isScramble2("abcde", "caebd"))
	assert.True(t, isScramble2("a", "a"))
}
