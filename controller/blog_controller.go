package controller

import (
	"blog/model"
	service "blog/service/blogservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBlog(context *gin.Context) {
	title := context.PostForm("title")
	content := context.PostForm("content")
	category := context.PostForm("category")
	author := context.PostForm("author")

	message, err := service.AddBlogService(model.DB, title, content, category, author)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}

func UpdateBlog(context *gin.Context){
	title := context.PostForm("title")
	content := context.PostForm("content")
	category := context.PostForm("category")
	author := context.PostForm("author")
	blog := context.PostForm("blog")

	message, err := service.UpdateBlogService(model.DB, title, content, category, author, blog)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}

func DeleteBlog(context *gin.Context){
	author := context.PostForm("author")
	blog := context.PostForm("blog")

	message, err := service.DeleteBlogService(model.DB, author, blog)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	context.IndentedJSON(201, message)
}


func GetGuestBlog(context *gin.Context){
	limit := context.Request.URL.Query().Get("_limit")
	page := context.Request.URL.Query().Get("_page")
	categoryName := context.Request.URL.Query().Get("category")
	blogName := context.Request.URL.Query().Get("blog")

	message, err := service.GetGuestBlogServcie(limit, page, blogName, categoryName)
	
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	context.IndentedJSON(200, message)
}