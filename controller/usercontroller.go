package controller


import(
	"blog/model"
	"blog/service"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SelectUser(context *gin.Context){
	limit := context.Request.URL.Query().Get("_limit")
	page := context.Request.URL.Query().Get("_page")
	userName := context.Request.URL.Query().Get("UName")

	viewUser, err := service.GetUserService(model.DB, userName, page, limit)

	if err != nil{
		context.IndentedJSON(http.StatusBadRequest, "get user failed: " + err.Error())
	}

	context.IndentedJSON(200, viewUser)
}