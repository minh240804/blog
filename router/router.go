package router

import (
	"blog/authorize"
	"blog/controller"

	//"blog/authorize"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/user", authorize.BasicAuthAdmin, controller.SelectUser)
	router.POST("/user", authorize.BasicAuthAdmin, controller.AddUser)
	router.Run("localhost:8080")

	return router
}
