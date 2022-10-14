package routes

import (
	"github.com/Lucasmartinsn/new-api-gin/handlers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		lol := main.Group("lol")
		{
			lol.GET("/", handlers.Get)
		}
	}

	return router
}