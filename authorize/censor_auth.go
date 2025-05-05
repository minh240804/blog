package authorize

import (
	model "blog/model"
	"blog/repository"

	"github.com/gin-gonic/gin"
)

func BasicAuthCensor(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	slicesAdmin, err := repository.GetCensor(model.DB)

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
