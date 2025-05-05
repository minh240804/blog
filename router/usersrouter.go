package router

import (
	"blog/authorize"
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	// user := rg.Group("")


	rg.GET("/", authorize.BasicAuthAdmin, controller.SelectUser)
	rg.POST("/", authorize.BasicAuthAdmin, controller.AddUser)
	rg.PUT("/", authorize.BasicAuthAdmin, controller.UppdateUser)
	rg.DELETE("/", authorize.BasicAuthAdmin, controller.DeleteUser)
	rg.PATCH("/", authorize.BasicAuthAdmin, controller.Decentralization)

}
