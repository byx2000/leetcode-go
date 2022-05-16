package best_time_to_buy_and_sell_stock_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 6, maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	assert.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
	assert.Equal(t, 0, maxProfit([]int{1}))
}
