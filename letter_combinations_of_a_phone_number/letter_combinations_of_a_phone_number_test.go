package letter_combinations_of_a_phone_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.Equal(t, []string{}, letterCombinations(""))
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
}
