package routers

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/pin-yu/datalab-drinks-backend/src/utils"
	cors "github.com/rs/cors/wrapper/gin"
)

// Run the web server
func Run() {
	router := setupRouter()
	router.RunTLS(
		":5002",
		filepath.Join(utils.GetBasePath(), "../../certs/fullchain.pem"),
		filepath.Join(utils.GetBasePath(), "../../certs/privkey.pem"),
	)
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
