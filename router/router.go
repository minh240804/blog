package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(config string) *gin.Engine {

	if config == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	v1 := router.Group("/user")
	addUserRoutes(v1)

	v2 := router.Group("/category")
	addCategoryRoutes(v2)

	v3 := router.Group("/blog")
	addBlogRoutes(*v3)
	return router
}
