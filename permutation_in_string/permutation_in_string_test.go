package permutation_in_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.True(t, checkInclusion("ab", "eidbaooo"))
	assert.False(t, checkInclusion("ab", "eidboaoo"))
	assert.False(t, checkInclusion("ab", "a"))
}
