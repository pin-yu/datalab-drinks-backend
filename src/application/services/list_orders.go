package services

import (
	"net/http"

	"github.com/pinyu/datalab-drinks-backend/src/infrastructure/persistence"
	"github.com/pinyu/datalab-drinks-backend/src/interface/responses"
)

// ListOrders returns week orders
func ListOrders() *responses.Response {
	// create a order repository to handle orders in database
	orderRepository := persistence.NewOrderRepository()

	orders, err := orderRepository.QueryWeekOrders()
	if err != nil {
		return responses.NewResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	// build response of week orders
	weekOrderResponse := responses.NewWeekOrdersResponse(orders)

	// put weekOrderResponse into basic response and return
	return responses.NewResponse(http.StatusAccepted, "ok", weekOrderResponse)
}
