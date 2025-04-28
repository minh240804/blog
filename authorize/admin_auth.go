package authorize

import (
	model "blog/model"
	"blog/repository"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Retrieve the token from the Authorization header
	token := c.GetHeader("Authorization")
	//log.Print("Token: ", token)
	if token != "secret_token" {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

func BasicAuthAdmin(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	slicesAdmin, err := repository.GetAdmin(model.DB)

	if err != nil {
		c.JSON(404, gin.H{"message": "data not found: " + err.Error()})
		c.Abort()
		return
	}

	for _, s := range slicesAdmin {
		if username == s.UserName && password == s.Password {
			c.Next()
			return
		}
	}
	c.JSON(401, gin.H{"message": "Unauthorized"})
	c.Abort()

}
