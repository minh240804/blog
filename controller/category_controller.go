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

func AddCategory(context *gin.Context) {
	category := context.PostForm("category")
	description := context.PostForm("description")

	message, err := service.AddCategoryService(model.DB, category, description)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}

func UpdateCategory(context *gin.Context) {
	category := context.PostForm("category")
	description := context.PostForm("description")
	id := context.PostForm("categoryId")

	message, err := service.UpdateCategory(model.DB, id, category, description)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}


func DeleteCategory(context *gin.Context) {
	id := context.Request.URL.Query().Get("id")

	message, err := service.DeleteCategoryService(model.DB, id)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}
