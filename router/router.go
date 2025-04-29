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
	router.PUT("/user", authorize.BasicAuthAdmin, controller.UppdateUser)
	router.DELETE("/user", authorize.BasicAuthAdmin, controller.DeleteUser)
	router.PATCH("/user", authorize.BasicAuthAdmin, controller.Decentralization)

	router.GET("/category", controller.SelectCategory)
	router.POST("/category", authorize.BasicAuthAdmin, controller.AddCategory)

	router.Run("localhost:8080")

	return router
}
