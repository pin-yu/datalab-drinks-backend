package services

import (
	"github.com/pinyu/datalab-drinks-backend/src/domain/orders"
)

// ListOrders returns week orders
func ListOrders() *orders.Orders {
	return orders.GetOrders()
}
