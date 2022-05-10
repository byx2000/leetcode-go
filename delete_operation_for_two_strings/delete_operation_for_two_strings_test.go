package delete_operation_for_two_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, minDistance("sea", "eat"))
	assert.Equal(t, 4, minDistance("leetcode", "etco"))
}
