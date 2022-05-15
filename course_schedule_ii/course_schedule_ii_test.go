package course_schedule_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Contains(t, [][]int{{0, 1}}, findOrder(2, [][]int{{1, 0}}))
	assert.Contains(t, [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}}, findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	assert.Contains(t, [][]int{{0}}, findOrder(1, [][]int{}))
}
