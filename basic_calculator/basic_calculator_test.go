package basic_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, 2, Calculate("1 + 1"))
	assert.Equal(t, 4, Calculate("6-4 / 2"))
	assert.Equal(t, 3, Calculate(" 2-1 + 2"))
	assert.Equal(t, 23, Calculate("(1+(4+5+2)-3)+(6+8)"))
	assert.Equal(t, 7, Calculate("3+2*2"))
}
