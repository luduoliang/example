package admin

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	var testData = map[string]interface{}{
		"id":   1,
		"name": "admin王五",
	}
	c.JSON(200, gin.H{
		"status": true,
		"msg":    "success",
		"data":   testData,
	})
}
