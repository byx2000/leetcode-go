package online_stock_span

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	s := Constructor()
	assert.Equal(t, 1, s.Next(100))
	assert.Equal(t, 1, s.Next(80))
	assert.Equal(t, 1, s.Next(60))
	assert.Equal(t, 2, s.Next(70))
	assert.Equal(t, 1, s.Next(60))
	assert.Equal(t, 4, s.Next(75))
	assert.Equal(t, 6, s.Next(85))
}
