package router

import (
	"blog/controller"
	//"blog/authorize"

	"github.com/gin-gonic/gin"
)
	
func InitRouter() *gin.Engine{
	router := gin.Default()

	//router.GET("/user", authorize.BasicAuthAdmin,controller.SelectUser)
	router.POST("/user", controller.AddUser)
	router.Run("localhost:8080")

	return router
}