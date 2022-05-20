package my_calendar_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	c := Constructor()
	assert.Equal(t, 1, c.Book(10, 20))
	assert.Equal(t, 1, c.Book(50, 60))
	assert.Equal(t, 2, c.Book(10, 40))
	assert.Equal(t, 3, c.Book(5, 15))
	assert.Equal(t, 3, c.Book(5, 10))
	assert.Equal(t, 3, c.Book(25, 55))
}
