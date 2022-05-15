package bus_routes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
	assert.Equal(t, -1, numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12))
	assert.Equal(t, 0, numBusesToDestination([][]int{{1, 7}, {3, 5}}, 5, 5))
	assert.Equal(t, 1, numBusesToDestination([][]int{{1, 9, 12, 20, 23, 24, 35, 38}, {10, 21, 24, 31, 32, 34, 37, 38, 43}, {10, 19, 28, 37}, {8}, {14, 19}, {11, 17, 23, 31, 41, 43, 44}, {21, 26, 29, 33}, {5, 11, 33, 41}, {4, 5, 8, 9, 24, 44}}, 37, 28))
	assert.Equal(t, 3, numBusesToDestination([][]int{{10, 13, 22, 28, 32, 35, 43}, {2, 11, 15, 25, 27}, {6, 13, 18, 25, 42}, {5, 6, 20, 27, 37, 47}, {7, 11, 19, 23, 35}, {7, 11, 17, 25, 31, 43, 46, 48}, {1, 4, 10, 16, 25, 26, 46}, {7, 11}, {3, 9, 19, 20, 21, 24, 32, 45, 46, 49}, {11, 41}}, 37, 43))
}

func Test2(t *testing.T) {
	assert.Equal(t, 2, numBusesToDestination2([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
	assert.Equal(t, -1, numBusesToDestination2([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12))
	assert.Equal(t, 0, numBusesToDestination2([][]int{{1, 7}, {3, 5}}, 5, 5))
	assert.Equal(t, 1, numBusesToDestination2([][]int{{1, 9, 12, 20, 23, 24, 35, 38}, {10, 21, 24, 31, 32, 34, 37, 38, 43}, {10, 19, 28, 37}, {8}, {14, 19}, {11, 17, 23, 31, 41, 43, 44}, {21, 26, 29, 33}, {5, 11, 33, 41}, {4, 5, 8, 9, 24, 44}}, 37, 28))
	assert.Equal(t, 3, numBusesToDestination2([][]int{{10, 13, 22, 28, 32, 35, 43}, {2, 11, 15, 25, 27}, {6, 13, 18, 25, 42}, {5, 6, 20, 27, 37, 47}, {7, 11, 19, 23, 35}, {7, 11, 17, 25, 31, 43, 46, 48}, {1, 4, 10, 16, 25, 26, 46}, {7, 11}, {3, 9, 19, 20, 21, 24, 32, 45, 46, 49}, {11, 41}}, 37, 43))
}
