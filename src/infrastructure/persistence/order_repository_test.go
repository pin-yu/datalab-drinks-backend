package persistence

import (
	"testing"
	"time"

	"github.com/pin-yu/datalab-drinks-backend/src/domain/entities"
	"github.com/stretchr/testify/assert"
)

// Only test time dependent functions
// e.g. ORDER BY > timestamp
func TestTimeDependentFunctions(t *testing.T) {
	setupTestDB()

	testHasOrder(t)
	testQueryOrders(t)

	cleanTestDB()
}

func testHasOrder(t *testing.T) {
	db := newDBDriver()

	// create an order with fake time stamp
	db.Create(&entities.Order{
		OrderBy:   "pinyu",
		Size:      "large",
		ItemID:    1,
		SugarID:   1,
		IceID:     1,
		OrderTime: time.Now().Unix(),
	})

	assert.False(t, NewOrderRepository().HasOrdered("pinyu"))

	// create an order with fake time stamp
	db.Create(&entities.Order{
		OrderBy:   "pinyu",
		Size:      "large",
		ItemID:    1,
		SugarID:   1,
		IceID:     1,
		OrderTime: time.Now().UnixNano(),
	})

	assert.True(t, NewOrderRepository().HasOrdered("pinyu"))
}

func testQueryOrders(t *testing.T) {
	db := newDBDriver()

	db.Create(&entities.Order{
		OrderBy:   "yilu",
		Size:      "large",
		ItemID:    1,
		SugarID:   1,
		IceID:     1,
		OrderTime: time.Now().Unix(),
	})

	db.Create(&entities.Order{
		OrderBy:   "yuqiao",
		Size:      "large",
		ItemID:    1,
		SugarID:   1,
		IceID:     1,
		OrderTime: time.Now().UnixNano(),
	})

	orders, _ := NewOrderRepository().QueryOrders()
	assert.Equal(t, 1, len(orders.AggregateOrders))
	assert.Equal(t, 2, len(orders.DetailOrders))
}
