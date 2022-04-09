package admin

import (
	"example/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//登录
func Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.Request.FormValue("name")
		password := c.Request.FormValue("pwd")
		if username == "" {
			c.JSON(200, gin.H{
				"status": false,
				"msg":    "请输入用户名",
				"data":   nil,
			})
			return
		}
		if password == "" {
			c.JSON(200, gin.H{
				"status": false,
				"msg":    "请输入密码",
				"data":   nil,
			})
			return
		}

		middleware.SetSession(c, "admin_info", "admin_info")
		c.JSON(200, gin.H{
			"status": true,
			"msg":    "登录成功",
			"data":   "admin_info",
		})
	} else {
		c.HTML(http.StatusOK, "admin/login/login.tmpl", nil)
	}
}

//退出登录
func Logout(c *gin.Context) {
	middleware.DelSession(c, "admin_info")
	c.Redirect(http.StatusFound, "/admin/login")
}
