package course_schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.True(t, canFinish(2, [][]int{{1, 0}}))
	assert.False(t, canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
