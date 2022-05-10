package largest_sum_of_averages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.InDelta(t, 20.00, largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3), 1e-6)
	assert.InDelta(t, 20.5, largestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4), 1e-6)
}
