package router

import (
	"blog/authorize"
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func addBlogRoutes(rg gin.RouterGroup) {

	rg.POST("/", authorize.BasicAuthLogin, controller.AddBlog)
	rg.PUT("/", authorize.BasicAuthLogin, controller.UpdateBlog)
	rg.DELETE("/", authorize.BasicAuthLogin, controller.DeleteBlog)
	rg.GET("/Guest/", controller.GetGuestBlog)
}