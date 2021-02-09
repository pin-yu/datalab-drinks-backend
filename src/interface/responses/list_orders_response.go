package responses

import "github.com/pinyu/datalab-drinks-backend/src/domain/entities"

// WeekOrdersResponse contains detail orders this week
type WeekOrdersResponse struct {
	DetailOrders []WeekOrderResponse `json:"week_orders"`
}

// WeekOrderResponse is a struct of single detail orders
type WeekOrderResponse struct {
	OrderBy  string `json:"order_by"`
	Item     string `json:"item"`
	Price    uint   `json:"price"`
	SugarTag string `json:"sugar_tag"`
	IceTag   string `json:"ice_tag"`
}

// NewWeekOrdersResponse retrieves order array and return a pointer of DetailOrders
func NewWeekOrdersResponse(orders []entities.Order) *WeekOrdersResponse {
	detailOrders := WeekOrdersResponse{}
	for _, order := range orders {
		detailOrders.DetailOrders = append(detailOrders.DetailOrders, WeekOrderResponse{
			OrderBy:  order.OrderBy,
			Item:     order.Item.Item,
			Price:    order.Price(),
			SugarTag: order.Sugar.Tag,
			IceTag:   order.Ice.Tag,
		})
	}
	return &detailOrders
}
