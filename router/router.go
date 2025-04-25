package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)
	
func InitRouter() *gin.Engine{
	router := gin.Default()

	router.GET("/user", controller.SelectUser)
	router.Run("localhost:8080")


	return router
}