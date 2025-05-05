package authorize

import "github.com/gin-gonic/gin"

func BasicAuthLogin(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	if username == "" || password == ""{
		c.JSON(401, gin.H{"message": "Unauthorized "})
		c.Abort()
		return
	}

	c.Next()
}
