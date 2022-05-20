package minimum_number_of_k_consecutive_bit_flips

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, minKBitFlips([]int{0, 1, 0}, 1))
	assert.Equal(t, -1, minKBitFlips([]int{1, 1, 0}, 2))
	assert.Equal(t, 3, minKBitFlips([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
}
