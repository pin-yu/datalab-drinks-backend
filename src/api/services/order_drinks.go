package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pinyu/datalab-drinks-backend/src/infra/orm"
)

// OrderDrinks returns a http result and a json message
func OrderDrinks(c *gin.Context) (int, string) {
	order := parseOrderQuery(c)
	orm.OrderExist(order)
	orm.WriteOrder(order)

	return http.StatusOK, "drinks are ordered"
}

func parseOrderQuery(c *gin.Context) *orm.Order {
	user := c.Query("order_by")
	itemID := parseUint8(c.Query("item_id"))
	sugar := parseUint8(c.Query("sugar"))
	ice := parseUint8(c.Query("ice"))

	order := &orm.Order{
		OrderBy: user,
		Item:    itemID,
		Sugar:   sugar,
		Ice:     ice,
	}
	return order
}

func parseUint8(s string) uint8 {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Printf("bad value in parseUint8: %v", s)
	}
	valueUint8 := uint8(value)

	return valueUint8
}
