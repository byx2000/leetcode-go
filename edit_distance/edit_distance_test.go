package edit_distance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, minDistance("horse", "ros"))
	assert.Equal(t, 5, minDistance("intention", "execution"))
}
