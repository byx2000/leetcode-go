package daily_temperatures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, dailyTemperatures1([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	assert.Equal(t, []int{1, 1, 1, 0}, dailyTemperatures1([]int{30, 40, 50, 60}))
	assert.Equal(t, []int{1, 1, 0}, dailyTemperatures1([]int{30, 60, 90}))
}

func Test2(t *testing.T) {
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, dailyTemperatures2([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	assert.Equal(t, []int{1, 1, 1, 0}, dailyTemperatures2([]int{30, 40, 50, 60}))
	assert.Equal(t, []int{1, 1, 0}, dailyTemperatures2([]int{30, 60, 90}))
}
