package car_pooling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.False(t, carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	assert.True(t, carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	assert.True(t, carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
	assert.False(t, carPooling([][]int{{9, 0, 1}, {3, 3, 7}}, 4))
}
