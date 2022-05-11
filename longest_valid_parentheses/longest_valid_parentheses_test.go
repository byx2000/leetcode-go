package longest_valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, longestValidParentheses("(()"))
	assert.Equal(t, 4, longestValidParentheses(")()())"))
}
