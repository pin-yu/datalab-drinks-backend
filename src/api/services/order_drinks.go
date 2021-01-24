package services

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pinyu/datalab-drinks-backend/src/infra/orm"
)

// OrderDrinks returns a http result and a json message
func OrderDrinks(c *gin.Context) string {
	orm.NewDB()

	user := c.Query("order_by")
	itemID := c.Query("item_id")
	sugar, err := strconv.ParseUint(c.Query("sugar"), 10, 64)
	if err != nil {
		log.Panic("bad sugar value")
	}
	sugarUint8 := uint8(sugar)

	ice, err := strconv.ParseUint(c.Query("ice"), 10, 64)
	if err != nil {
		log.Panic("bad ice value")
	}
	iceUint8 := uint8(ice)

	order := orm.Order{
		OrderBy: user,
		Item:    itemID,
		Sugar:   sugarUint8,
		Ice:     iceUint8,
	}

	fmt.Printf("user: %v, itemId: %v, sugar: %v, ice: %v", user, itemID, sugar, ice)
	return "put orders"

	// write to db
}
