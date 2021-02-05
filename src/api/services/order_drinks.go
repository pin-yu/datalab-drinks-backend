package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pinyu/datalab-drinks-backend/src/infrastructure/orm"
	"github.com/pinyu/datalab-drinks-backend/src/infrastructure/orm/schemas"
)

// OrderDrinks returns a http result and a json message
func OrderDrinks(c *gin.Context) (int, string) {
	order := parseOrderQuery(c)
	orm.OrderExist(order)
	orm.WriteOrder(order)

	return http.StatusOK, "drinks are ordered"
}

func parseOrderQuery(c *gin.Context) *schemas.Order {
	var result struct {
		OrderBy string `json:"order_by"`
		ItemID  uint   `json:"item_id"`
		Sugar   uint   `json:"sugar"`
		Ice     uint   `json:"ice"`
	}
	c.Bind(&result)

	order := &schemas.Order{
		OrderBy: result.OrderBy,
		Item:    uint8(result.ItemID),
		Sugar:   uint8(result.Sugar),
		Ice:     uint8(result.Ice),
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
