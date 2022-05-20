package palindromic_substrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, countSubstrings("abc"))
	assert.Equal(t, 6, countSubstrings("aaa"))
}
