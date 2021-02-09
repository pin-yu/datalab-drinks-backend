package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pinyu/datalab-drinks-backend/src/application/services"
)

func addMenusRoutes(rg *gin.RouterGroup) {
	drinks := rg.Group("/menus")

	drinks.GET("/", func(c *gin.Context) {
		res := services.ReadCamaMenu()
		c.JSON(res.Resolve())
	})
}
