package target_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 5, findTargetSumWays1([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(t, 1, findTargetSumWays1([]int{1}, 1))
	assert.Equal(t, 256, findTargetSumWays1([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}

func Test2(t *testing.T) {
	assert.Equal(t, 5, findTargetSumWays2([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(t, 1, findTargetSumWays2([]int{1}, 1))
	assert.Equal(t, 256, findTargetSumWays2([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}

func Test3(t *testing.T) {
	assert.Equal(t, 5, findTargetSumWays3([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(t, 1, findTargetSumWays3([]int{1}, 1))
	assert.Equal(t, 256, findTargetSumWays3([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
