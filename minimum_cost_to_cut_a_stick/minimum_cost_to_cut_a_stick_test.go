package minimum_cost_to_cut_a_stick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 16, minCost(7, []int{1, 3, 4, 5}))
	assert.Equal(t, 22, minCost(9, []int{5, 6, 1, 4, 2}))
}
