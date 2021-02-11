package entities

import "time"

// AggregateOrder represents the summation of week's orders
type AggregateOrder struct {
	ItemID        uint
	Item          string
	Size          string
	SugarTag      string
	IceTag        string
	SubTotalPrice uint
	Number        uint
}

// Order represents basic drinks orders
type Order struct {
	ID        uint   `gorm:"primaryKey"`
	OrderBy   string `gorm:"not null"`
	ItemID    uint
	Item      Item
	Size      string
	SugarID   uint
	Sugar     Sugar
	IceID     uint
	Ice       Ice
	OrderTime int64
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

// OrderTimeToRFC3339 returns order time in RFC3339 format
func (o *Order) OrderTimeToRFC3339() string {
	return time.Unix(0, o.OrderTime).Format(time.RFC3339)
}
