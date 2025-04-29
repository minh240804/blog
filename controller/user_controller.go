package controller

import (
	"blog/model"
	userService "blog/service/userservice"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SelectUser(context *gin.Context) {
	limit := context.Request.URL.Query().Get("_limit")
	page := context.Request.URL.Query().Get("_page")
	userName := context.Request.URL.Query().Get("username")
	role := context.Request.URL.Query().Get("role")

	viewUser, err := userService.GetUserService(model.DB, userName, limit, page, role)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, "get user failed: "+err.Error())
		return
	}

	context.IndentedJSON(200, viewUser)
}

func AddUser(context *gin.Context) {
	userName := context.PostForm("username")
	password := context.PostForm("password")
	role := context.PostForm("role")

	message, err := userService.AddUserService(model.DB, userName, password, role)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}

func UppdateUser(context *gin.Context) {
	userName := context.PostForm("username")
	password := context.PostForm("password")
	role := context.PostForm("role")
	id := context.PostForm("id")

	message, err := userService.UpdateUserService(model.DB, userName, password, role, id)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}


func DeleteUser(context *gin.Context) {
	id := context.Request.URL.Query().Get("id")
	message, err := userService.DeleteUserService(model.DB, id)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}

func Decentralization(context *gin.Context) {
	id := context.Request.URL.Query().Get("id")
	message, err := userService.DeleteUserService(model.DB, id)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(201, message)
}
