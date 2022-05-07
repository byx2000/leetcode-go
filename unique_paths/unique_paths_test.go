package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 28, uniquePaths(3, 7))
	assert.Equal(t, 3, uniquePaths(3, 2))
	assert.Equal(t, 28, uniquePaths(7, 3))
	assert.Equal(t, 6, uniquePaths(3, 3))
}
