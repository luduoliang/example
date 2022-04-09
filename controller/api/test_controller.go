package api

import (
	"example/middleware"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	username, exists := c.Get("username")
	usernameStr := ""
	if exists {
		usernameStr = username.(string)
	}
	var testData = map[string]interface{}{
		"id":   1,
		"name": "王五" + usernameStr,
	}
	c.JSON(200, gin.H{
		"status": true,
		"msg":    "success",
		"data":   testData,
	})
}

func GetJwtToken(c *gin.Context) {
	token, _ := middleware.GenToken("admin")
	c.JSON(200, gin.H{
		"status": true,
		"msg":    "success",
		"data":   token,
	})
}
