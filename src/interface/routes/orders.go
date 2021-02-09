package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pinyu/datalab-drinks-backend/src/application/services"
)

func addOrdersRoutes(rg *gin.RouterGroup) {
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
