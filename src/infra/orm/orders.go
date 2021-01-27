package orm

import (
	"fmt"
	"log"
	"time"

	"github.com/pinyu/datalab-drinks-backend/src/utils"
)

/*
	NOTICE: this package should not contain the domain logic
	TODO: create Item, Sugar, Ice tables
*/

// Order will be pluralized to Orders by gorm
type Order struct {
	ID        uint
	OrderBy   string
	Item      uint8
	Sugar     uint8
	Ice       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o Order) String() string {
	return fmt.Sprintf("{ID: %v, OrderBy: %v, Item: %v, Sugar: %v, Ice: %v, CreatedAt: %v, UpdatedAt: %v}",
		o.ID,
		o.OrderBy,
		o.Item,
		o.Sugar,
		o.Ice,
		o.CreatedAt,
		o.UpdatedAt,
	)
}

func migrateOrder() {
	db := newDB()
	db.AutoMigrate(&Order{})
	log.Println("migrate order table")
}

func dropOrder() {
	db := newDB()
	db.Migrator().DropTable(&Order{})
	log.Println("drop order table")
}

// service should not use orm's method
// I should move the logic to the domain package

// OrderExist returns true if the specified order exists
func OrderExist(order *Order) bool {
	db := newDB()

	theOrder := Order{}
	result := db.Where("order_by = ? AND created_at > ?", order.OrderBy, utils.LastFridayNoon()).Order("created_at desc").First(&theOrder)

	// orders exist
	if result.RowsAffected > 0 {
		log.Printf("orders exist: %v", theOrder)
		return true
	}

	return false
}

// WriteOrder writes an order record into database
func WriteOrder(order *Order) {
	db := newDB()

	log.Println(order)

	result := db.Create(order)
	if result.Error != nil {
		log.Printf("error occurs in WriteOrder: %v", result.Error)
	}
}
