package coin_change_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 4, change(5, []int{1, 2, 5}))
	assert.Equal(t, 0, change(3, []int{2}))
	assert.Equal(t, 1, change(10, []int{10}))
}
