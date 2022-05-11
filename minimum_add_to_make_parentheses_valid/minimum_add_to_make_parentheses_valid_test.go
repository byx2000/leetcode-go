package minimum_add_to_make_parentheses_valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, minAddToMakeValid("())"))
	assert.Equal(t, 3, minAddToMakeValid("((("))
	assert.Equal(t, 4, minAddToMakeValid("()))(("))
}
