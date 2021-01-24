package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinyu/datalab-drinks-backend/src/api/services"
)

func addMenusRoutes(rg *gin.RouterGroup) {
	drinks := rg.Group("/menus")

	drinks.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, services.GetCamaMenus())
	})
}
