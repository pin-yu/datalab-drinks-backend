package services

import (
	"net/http"

	"github.com/pin-yu/datalab-drinks-backend/src/infrastructure/persistence"
	"github.com/pin-yu/datalab-drinks-backend/src/interface/responses"
)

// ListOrders returns week orders
func ListOrders() *responses.Response {
	// create a order repository to handle orders in database
	orderRepository := persistence.NewOrderRepository()

	orders, err := orderRepository.QueryOrders()
	if err != nil {
		return responses.NewResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	// build response of week orders
	weekOrderResponse := responses.NewWeekOrdersResponse(orders)

	// put weekOrderResponse into basic response and return
	return responses.NewResponse(http.StatusOK, "ok", weekOrderResponse)
}
