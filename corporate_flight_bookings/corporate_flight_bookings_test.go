package corporate_flight_bookings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{10, 55, 45, 25, 25}, corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
	assert.Equal(t, []int{10, 25}, corpFlightBookings([][]int{{1, 2, 10}, {2, 2, 15}}, 2))
}
