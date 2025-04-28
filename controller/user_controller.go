package controller

import (
	"blog/model"
	"blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SelectUser(context *gin.Context) {
	limit := context.Request.URL.Query().Get("_limit")
	page := context.Request.URL.Query().Get("_page")
	userName := context.Request.URL.Query().Get("UName")

	viewUser, err := service.GetUserService(model.DB, userName, limit, page)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, "get user failed: "+err.Error())
		return
	}

	context.IndentedJSON(201, viewUser)
}

func AddUser(context *gin.Context) {
	userName := context.PostForm("UserName")
	password := context.PostForm("Password")
	role := context.PostForm("Role")

	message, err := service.AddUserService(model.DB, userName, password, role)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(200, message)
}
