package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// Run the web server
func Run() {
	router := setupRouter()
	router.RunTLS(":5000", "./certs/server.crt", "./certs/server.key")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// support cors
	router.Use(cors.Default())

	v1 := router.Group("v1")
	addMenusRouter(v1)
	addOrdersRouter(v1)

	return router
}
