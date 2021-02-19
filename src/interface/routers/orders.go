package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pin-yu/datalab-drinks-backend/src/application/services"
)

func addOrdersRouter(rg *gin.RouterGroup) {
	orders := rg.Group("/orders")

	orders.GET("/", func(c *gin.Context) {
		res := services.ListOrders()
		c.JSON(res.Resolve())
	})

	orders.POST("/", func(c *gin.Context) {
		res := services.OrderDrinks(c)
		c.JSON(res.Resolve())
	})
}
