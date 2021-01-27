package routes

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

var (
	router = gin.Default()
)

// Run the web server
func Run() {
	router.Use(cors.Default())
	getRoutes()
	router.Run(":5000")
}

func getRoutes() {
	v1 := router.Group("v1")
	addMenusRoutes(v1)
	addUsersRoutes(v1)
	addOrdersRoutes(v1)
}
