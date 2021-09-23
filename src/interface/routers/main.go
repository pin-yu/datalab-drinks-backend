package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// Run the web server
func Run() {
	router := setupRouter()
	router.Run(":5002")
	// router.RunTLS(
	// 	":5002",
	// 	filepath.Join(utils.GetBasePath(), "../../certs/fullchain.pem"),
	// 	filepath.Join(utils.GetBasePath(), "../../certs/privkey.pem"),
	// )
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// support cors
	router.Use(cors.Default())

	group := router.Group("v2")
	addMenusRouter(group)
	addOrdersRouter(group)

	return router
}
