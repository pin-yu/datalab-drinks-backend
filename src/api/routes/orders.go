package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinyu/datalab-drinks-backend/src/api/services"
)

func addOrdersRoutes(rg *gin.RouterGroup) {
	orders := rg.Group("/orders")

	orders.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "get orders")
	})

	orders.POST("/", func(c *gin.Context) {
		c.JSON(services.OrderDrinks(c))
	})
}
