package router

import (
	"blog/authorize"
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func addCategoryRoutes(rg *gin.RouterGroup) {

	rg.GET("/", controller.SelectCategory)
	rg.POST("/", authorize.BasicAuthAdmin, controller.AddCategory)
	rg.PUT("/", authorize.BasicAuthAdmin, controller.UpdateCategory)
	rg.DELETE("/", authorize.BasicAuthAdmin, controller.DeleteCategory)
}
