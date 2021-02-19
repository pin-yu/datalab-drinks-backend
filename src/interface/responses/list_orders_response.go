package responses

import "github.com/pin-yu/datalab-drinks-backend/src/domain/entities"

// OrdersResponse represents the response of list_orders service
type OrdersResponse struct {
	MeetingTime     string                   `json:"meeting_time"`
	TotalPrice      uint                     `json:"total_price"`
	AggregateOrders []AggregateOrderResponse `json:"aggregate_orders"`
	DetailOrders    []OrderResponse          `json:"week_orders"`
}

// AggregateOrderResponse represents athe aggregate order
type AggregateOrderResponse struct {
	Item          string `json:"item"`
	Size          string `json:"size"`
	SugarTag      string `json:"sugar_tag"`
	IceTag        string `json:"ice_tag"`
	SubTotalPrice uint   `json:"sub_total_price"`
	Number        uint   `json:"number"`
}

// OrderResponse represents the detail order
type OrderResponse struct {
	OrderBy   string `json:"order_by"`
	Item      string `json:"item"`
	Size      string `json:"size"`
	Price     uint   `json:"price"`
	SugarTag  string `json:"sugar_tag"`
	IceTag    string `json:"ice_tag"`
	OrderTime string `json:"order_time"`
}

// NewWeekOrdersResponse retrieves order array and return a pointer of DetailOrders
func NewWeekOrdersResponse(orders *entities.Orders) *OrdersResponse {
	ordersResponse := OrdersResponse{}

	detailOrdersToResponse(&ordersResponse, orders)
	aggregateOrdersToResponse(&ordersResponse, orders)

	ordersResponse.MeetingTime = orders.MeetingTime
	ordersResponse.TotalPrice = orders.TotalPrice

	return &ordersResponse
}

func detailOrdersToResponse(ordersResponse *OrdersResponse, orders *entities.Orders) {
	for _, order := range orders.DetailOrders {
		ordersResponse.DetailOrders = append(ordersResponse.DetailOrders, OrderResponse{
			OrderBy:   order.OrderBy,
			Size:      order.Size,
			Item:      order.Item.Item,
			Price:     order.Price(),
			SugarTag:  order.Sugar.Tag,
			IceTag:    order.Ice.Tag,
			OrderTime: order.OrderTimeToRFC3339(),
		})
	}
}

func aggregateOrdersToResponse(ordersResponse *OrdersResponse, orders *entities.Orders) {
	for _, aggOrder := range orders.AggregateOrders {
		ordersResponse.AggregateOrders = append(ordersResponse.AggregateOrders, AggregateOrderResponse{
			Item:          aggOrder.Item,
			Size:          aggOrder.Size,
			SugarTag:      aggOrder.SugarTag,
			IceTag:        aggOrder.IceTag,
			SubTotalPrice: aggOrder.SubTotalPrice,
			Number:        aggOrder.Number,
		})
	}
}
