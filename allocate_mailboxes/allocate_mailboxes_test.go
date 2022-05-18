package allocate_mailboxes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 5, minDistance([]int{1, 4, 8, 10, 20}, 3))
	assert.Equal(t, 9, minDistance([]int{2, 3, 5, 12, 18}, 2))
	assert.Equal(t, 8, minDistance([]int{7, 4, 6, 1}, 1))
	assert.Equal(t, 0, minDistance([]int{3, 6, 14, 10}, 4))
}
