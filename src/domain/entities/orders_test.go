package entities

import (
	"testing"
	"time"

	"github.com/pin-yu/datalab-drinks-backend/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestCountTotalPrice(t *testing.T) {
	aggregateOrders := []AggregateOrder{
		{
			SubTotalPrice: 100,
		},
		{
			SubTotalPrice: 1000,
		},
	}

	orders := Orders{
		AggregateOrders: aggregateOrders,
	}

	orders.CountTotalPrice()

	assert.Equal(t, uint(1100), orders.TotalPrice)
}

func TestNewOrders(t *testing.T) {
	orders := NewOrders()

	assert.Equal(t, utils.MeetingStartTime().Format(time.RFC3339), orders.MeetingTime)
}
