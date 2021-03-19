package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pin-yu/datalab-drinks-backend/src/domain/repositories"
	"github.com/pin-yu/datalab-drinks-backend/src/infrastructure/persistence"
	"github.com/pin-yu/datalab-drinks-backend/src/interface/requests"
	"github.com/pin-yu/datalab-drinks-backend/src/interface/responses"
)

// OrderDrinks will validate request body and save order to database
func OrderDrinks(c *gin.Context) *responses.Response {
	// 1. validate request body
	orderRequest, err := parseOrderBody(c)
	if err != nil {
		return responses.NewResponse(http.StatusBadRequest, err.Error(), nil)
	}

	// create a order repository to handle orders in database
	orderRepository := persistence.NewOrderRepository()

	// 2. validate order request
	if err := validateOrderRequest(orderRequest, orderRepository); err != nil {
		return responses.NewResponse(http.StatusBadRequest, err.Error(), nil)
	}

	// 3. check if the order exists by checking order_by field
	statusMsg := orderExist(orderRepository, orderRequest.OrderBy)

	// 4. create/update an order record
	if err := orderRepository.CreateOrder(orderRequest); err != nil {
		return responses.NewResponse(http.StatusBadRequest, err.Error(), nil)
	}

	return responses.NewResponse(http.StatusOK, statusMsg, nil)
}

func parseOrderBody(c *gin.Context) (*requests.OrderRequestBody, error) {
	orderRequest := requests.OrderRequestBody{}

	err := c.Bind(&orderRequest)
	if err != nil || !orderRequest.IsSchemaValid() {
		return nil, fmt.Errorf("bad schema of order request body")
	}

	return &orderRequest, nil
}

func orderExist(orderRepository repositories.OrderRepository, orderBy string) string {
	var statusMsg string

	if orderRepository.HasOrdered(orderBy) {
		statusMsg = "drinks has been updated"
	} else {
		statusMsg = "drinks has been ordered"
	}

	return statusMsg
}

func validateOrderRequest(orderRequest *requests.OrderRequestBody, orderRepository repositories.OrderRepository) error {
	item, err := orderRepository.ValidateItemID(orderRequest.ItemID)
	if err != nil {
		return err
	}

	sugar, err := orderRepository.ValidateSugarID(orderRequest.SugarID)
	if err != nil {
		return err
	}

	ice, err := orderRepository.ValidateIceID(orderRequest.IceID)
	if err != nil {
		return err
	}

	if !item.Sugar && !sugar.IsNormalSugar() {
		return fmt.Errorf("the sugar should be normal")
	}

	// if a drinks cannot be made as hot, iceId couldn't be 1 which is hot
	if !item.Hot && ice.IsHot() {
		return fmt.Errorf("the drinks should be ice")
	}
	if !item.Cold && !ice.IsHot() {
		return fmt.Errorf("the drinks should be hot")
	}

	return nil
}
