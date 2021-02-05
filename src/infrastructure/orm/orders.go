package orm

import (
	"log"

	"github.com/pinyu/datalab-drinks-backend/src/infrastructure/orm/schemas"
	"github.com/pinyu/datalab-drinks-backend/src/utils"
)

/*
	NOTICE: this package should not contain the domain logic
	TODO: create Item, Sugar, Ice tables
*/

func migrateOrder() {
	conn := newDBConnection()
	conn.AutoMigrate(&schemas.Order{})
	log.Println("migrate order table")
}

func dropOrder() {
	conn := newDBConnection()
	conn.Migrator().DropTable(&schemas.Order{})
	log.Println("drop order table")
}

// service should not use orm's method
// I should move the logic to the domain package

// TODO: price
// WeekOrder is
type WeekOrder struct {
	OrderBy   string
	Item      uint8
	Size      string
	Sugar     uint8
	Ice       uint8
	UpdatedAt string
}

// GetWeekOrders select order_by, item, size, ice, sugar, MAX(updated_at) from orders where updated_at > "2021-01-30 13-00-00" group by order_by;
func GetWeekOrders() *[]WeekOrder {
	db := newDBConnection()

	weekOrders := &[]WeekOrder{}

	db.Debug().Table("orders").Select("order_by, item, size, sugar, ice, Max(updated_at) as updated_at").Where("updated_at > ?", utils.OrderIntervalStartTime()).Group("order_by").Find(weekOrders)

	return weekOrders
}

// OrderExist returns true if the specified order exists
func OrderExist(order *schemas.Order) bool {
	db := newDBConnection()

	theOrder := schemas.Order{}
	result := db.Where("order_by = ? AND created_at > ?", order.OrderBy, utils.OrderIntervalStartTime()).Order("created_at desc").First(&theOrder)

	// orders exist
	if result.RowsAffected > 0 {
		log.Printf("orders exist: %v", theOrder)
		return true
	}

	return false
}

// WriteOrder writes an order record into database
func WriteOrder(order *schemas.Order) {
	db := newDBConnection()

	log.Println(order)

	// TODO: remove this code
	if order.Size == "" {
		order.Size = "medium"
	}

	result := db.Create(order)
	if result.Error != nil {
		log.Printf("error occurs in WriteOrder: %v", result.Error)
	}
}
