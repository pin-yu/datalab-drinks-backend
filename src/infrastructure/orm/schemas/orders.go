package schemas

import (
	"fmt"
	"time"
)

// Order will be pluralized to Orders by gorm
type Order struct {
	ID        uint
	OrderBy   string
	Item      uint8
	Size      string
	Sugar     uint8
	Ice       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o Order) String() string {
	return fmt.Sprintf("{ID: %v, OrderBy: %v, Item: %v, Size: %v, Sugar: %v, Ice: %v, CreatedAt: %v, UpdatedAt: %v}",
		o.ID,
		o.OrderBy,
		o.Item,
		o.Size,
		o.Sugar,
		o.Ice,
		o.CreatedAt,
		o.UpdatedAt,
	)
}
