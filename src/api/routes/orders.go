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

	// the order requests should be idempotent, use PUT instead of Post
	orders.PUT("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, services.OrderDrinks(c))
	})

	orders.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "delete orders")
	})
}
