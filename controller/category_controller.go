package controller

import (
	"blog/model"
	service "blog/service/categoryservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SelectCategory(context *gin.Context) {
	limit := context.Request.URL.Query().Get("_limit")
	page := context.Request.URL.Query().Get("_page")
	categoryName := context.Request.URL.Query().Get("category")

	viewUser, err := service.GetCategoryService(model.DB, limit, page, categoryName)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, "get user failed: "+err.Error())
		return
	}

	context.IndentedJSON(200, viewUser)
}

func AddCategory(context *gin.Context){
	category := context.PostForm("category")
	description := context.PostForm("description")

	message, err := service.AddCategoryService(model.DB, category, description)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}