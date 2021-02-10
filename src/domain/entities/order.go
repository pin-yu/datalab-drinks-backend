package entities

import "time"

// Order represents drinks orders
type Order struct {
	ID        uint   `json:"order_id" gorm:"primaryKey"`
	OrderBy   string `json:"order_by" gorm:"not null"`
	ItemID    uint   `json:"item_id"`
	Item      Item   `json:"item"`
	Size      string `json:"size"`
	SugarID   uint   `json:"sugar_id"`
	Sugar     Sugar  `json:"sugar"`
	IceID     uint   `json:"ice_id"`
	Ice       Ice    `json:"ice"`
	OrderTime int64  `json:"order_time"`
}

// Price returns price of sizes
func (o *Order) Price() uint {
	switch o.Size {
	case "medium":
		return o.Item.MediumPrice
	case "large":
		return o.Item.LargePrice
	default:
		return 0
	}
}

// RFC3339 returns order time in RFC3339 format
func (o *Order) RFC3339() string {
	return time.Unix(0, o.OrderTime).Format(time.RFC3339)
}
