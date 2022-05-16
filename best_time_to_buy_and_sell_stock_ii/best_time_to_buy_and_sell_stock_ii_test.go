package best_time_to_buy_and_sell_stock_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 7, maxProfit1([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 4, maxProfit1([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfit1([]int{7, 6, 4, 3, 1}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 7, maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 4, maxProfit2([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfit2([]int{7, 6, 4, 3, 1}))
}
