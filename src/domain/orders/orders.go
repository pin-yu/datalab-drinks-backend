package orders

import (
	"time"

	"github.com/pinyu/datalab-drinks-backend/src/domain/menus"
	"github.com/pinyu/datalab-drinks-backend/src/infrastructure/orm"
	"github.com/pinyu/datalab-drinks-backend/src/utils"
)

// Orders .
type Orders struct {
	OrderDate    time.Time     `json:"order_date"`
	DetailOrders []DetailOrder `json:"detail_orders"`
}

// DetailOrder .
type DetailOrder struct {
	OrderBy   string `json:"order_by"`
	OrderTime string `json:"order_time"`
	Item      string `json:"item"`
	Size      string `json:"size"`
	Sugar     string `json:"sugar"`
	Ice       string `json:"ice"`
}

// OrderKey .
type OrderKey struct {
	Item  uint8  `json:"item"`
	Size  string `json:"size"`
	Sugar string `json:"sugar"`
	Ice   string `json:"ice"`
}

// AggregateOrder .
type AggregateOrder struct {
	Item   string `json:"item"`
	Size   string `json:"size"`
	Sugar  string `json:"sugar"`
	Ice    string `json:"ice"`
	Number uint   `json:"number"`
}

// put the domain logic here please
func isValidOrder() {
	return
}

// GetOrders .
func GetOrders() *Orders {
	weekOrders := orm.GetWeekOrders()

	detailOrders := make([]DetailOrder, len(*weekOrders))

	for i, v := range *weekOrders {
		detailOrders[i] = *toDetailOrder(&v)
	}

	return &Orders{
		OrderDate:    utils.OrderIntervalEndTime(),
		DetailOrders: detailOrders,
	}
}

var itemMap = menus.GetItemMap()

func toDetailOrder(weekOrder *orm.WeekOrder) *DetailOrder {
	return &DetailOrder{
		OrderBy:   weekOrder.OrderBy,
		OrderTime: weekOrder.UpdatedAt,
		Item:      itemMap[weekOrder.Item],
		Size:      weekOrder.Size,
		Sugar:     sugarIDToLabel(weekOrder.Sugar),
		Ice:       iceIDToLabel(weekOrder.Ice),
	}
}

// func toOrderKey(detailOrder *DetailOrder) *OrderKey {
// 	return &OrderKey{
// 		Item: detailOrder.Item,
// 		Size: detailOrder.Size,
// 		Sugar: detailOrder.

// 	}
// }

func sugarIDToLabel(ID uint8) string {
	switch ID {
	case 1:
		return "無糖"
	case 2:
		return "微糖"
	case 3:
		return "少糖"
	case 4:
		return "正常糖"
	default:
		return "bad sugar id"
	}
}

func iceIDToLabel(ID uint8) string {
	switch ID {
	case 1:
		return "熱"
	case 2:
		return "少冰"
	case 3:
		return "正常冰"
	default:
		return "bad ice id"
	}
}

func getWeekOrders() {

}

func getCompactOrders() {

}
