package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
}
