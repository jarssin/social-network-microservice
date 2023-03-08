package server

import "github.com/gin-gonic/gin"

func SetupRoutes(engine *gin.Engine) {
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
