package entities

import (
	"time"

	"github.com/pin-yu/datalab-drinks-backend/src/utils"
)

// Orders represents the aggregation of orders
type Orders struct {
	MeetingTime     string
	TotalPrice      uint
	AggregateOrders []AggregateOrder
	DetailOrders    []Order
}

// NewOrders returns Orders with meeting time
func NewOrders() *Orders {
	orders := Orders{
		MeetingTime: utils.MeetingStartTime().Format(time.RFC3339),
	}

	return &orders
}

// CountTotalPrice will set the total price according to the AggregateOrders
func (o *Orders) CountTotalPrice() {
	totalPrice := uint(0)
	for _, aggOrder := range o.AggregateOrders {
		totalPrice += aggOrder.SubTotalPrice
	}

	o.TotalPrice = totalPrice
}
